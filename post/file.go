package post

import (
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	storage "google.golang.org/api/storage/v1"
)

const (
	// Change these variable to match your personal information.
	bucketName = "YOUR_BUCKET_NAME"
	projectID  = "YOUR_PROJECT_ID"

	fileName   = "/usr/share/dict/words" // The name of the local file to upload.
	objectName = "english-dictionary"    // This can be changed to any valid object name.

	// For the basic sample, these variables need not be changed.
	scope      = storage.DevstorageFull_controlScope
	entityName = "allUsers"
)

var (
	jsonFile = flag.String("creds", "", "A path to your JSON key file for your service account downloaded from Google Developer Console, not needed if you run it on Compute Engine instances.")
)

func fatalf(service *storage.Service, errorMessage string, args ...interface{}) {
	restoreOriginalState(service)
	log.Fatalf("Dying with error:\n"+errorMessage, args...)
}

func restoreOriginalState(service *storage.Service) bool {
	// Delete an object from a bucket.
	if err := service.Objects.Delete(bucketName, objectName).Do(); err != nil {
		// If the object exists but wasn't deleted, the bucket deletion will also fail.
		fmt.Printf("Could not delete object during cleanup: %v\n\n", err)
	} else {
		fmt.Printf("Successfully deleted %s/%s during cleanup.\n\n", bucketName, objectName)
	}

	// Delete a bucket in the project
	if err := service.Buckets.Delete(bucketName).Do(); err != nil {
		fmt.Printf("Could not delete bucket during cleanup: %v\n\n", err)
		fmt.Println("WARNING: Final cleanup attempt failed. Original state could not be restored.\n")
		return false
	}

	fmt.Printf("Successfully deleted bucket %s during cleanup.\n\n", bucketName)
	return true
}

func main() {
	flag.Parse()
	if *jsonFile != "" {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", *jsonFile)
	}

	client, err := google.DefaultClient(context.Background(), scope)
	if err != nil {
		log.Fatalf("Unable to get default client: %v", err)
	}
	service, err := storage.New(client)
	if err != nil {
		log.Fatalf("Unable to create storage service: %v", err)
	}

	// If the bucket already exists and the user has access, warn the user, but don't try to create it.
	if _, err := service.Buckets.Get(bucketName).Do(); err == nil {
		fmt.Printf("Bucket %s already exists - skipping buckets.insert call.", bucketName)
	} else {
		// Create a bucket.
		if res, err := service.Buckets.Insert(projectID, &storage.Bucket{Name: bucketName}).Do(); err == nil {
			fmt.Printf("Created bucket %v at location %v\n\n", res.Name, res.SelfLink)
		} else {
			fatalf(service, "Failed creating bucket %s: %v", bucketName, err)
		}
	}

	// List all buckets in a project.
	if res, err := service.Buckets.List(projectID).Do(); err == nil {
		fmt.Println("Buckets:")
		for _, item := range res.Items {
			fmt.Println(item.Id)
		}
		fmt.Println()
	} else {
		fatalf(service, "Buckets.List failed: %v", err)
	}

	// Insert an object into a bucket.
	object := &storage.Object{Name: objectName}
	file, err := os.Open(fileName)
	if err != nil {
		fatalf(service, "Error opening %q: %v", fileName, err)
	}
	if res, err := service.Objects.Insert(bucketName, object).Media(file).Do(); err == nil {
		fmt.Printf("Created object %v at location %v\n\n", res.Name, res.SelfLink)
	} else {
		fatalf(service, "Objects.Insert failed: %v", err)
	}

	// List all objects in a bucket
	if res, err := service.Objects.List(bucketName).Do(); err == nil {
		fmt.Printf("Objects in bucket %v:\n", bucketName)
		for _, object := range res.Items {
			fmt.Println(object.Name)
		}
		fmt.Println()
	} else {
		fatalf(service, "Objects.List failed: %v", err)
	}

	// Insert ACL for an object.
	// This illustrates the minimum requirements.
	objectAcl := &storage.ObjectAccessControl{
		Bucket: bucketName, Entity: entityName, Object: objectName, Role: "READER",
	}
	if res, err := service.ObjectAccessControls.Insert(bucketName, objectName, objectAcl).Do(); err == nil {
		fmt.Printf("Result of inserting ACL for %v/%v:\n%v\n\n", bucketName, objectName, res)
	} else {
		fatalf(service, "Failed to insert ACL for %s/%s: %v.", bucketName, objectName, err)
	}

	// Get ACL for an object.
	if res, err := service.ObjectAccessControls.Get(bucketName, objectName, entityName).Do(); err == nil {
		fmt.Printf("Users in group %v can access %v/%v as %v.\n\n",
			res.Entity, bucketName, objectName, res.Role)
	} else {
		fatalf(service, "Failed to get ACL for %s/%s: %v.", bucketName, objectName, err)
	}

	// Get an object from a bucket.
	if res, err := service.Objects.Get(bucketName, objectName).Do(); err == nil {
		fmt.Printf("The media download link for %v/%v is %v.\n\n", bucketName, res.Name, res.MediaLink)
	} else {
		fatalf(service, "Failed to get %s/%s: %s.", bucketName, objectName, err)
	}

	if !restoreOriginalState(service) {
		os.Exit(1)
	}
}
