import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"

func init() {
	awsAuth.LoadConfig()
}

    func TestDescribetables(t *testing.T) {

      var dynamodb Dynamodb

      describetables := map[string]interface{}{
        "Region":    "us-east-2",
        "TableName": "hello",
      }

    	_, err := bigtable.Describetables(describetables)

    	if err != nil {
    		t.Errorf("Test Fail")
    	}
    }



		func TestListtables(t *testing.T) {

			var dynamodb Dynamodb

			listtables := map[string]interface{}{
				"Region":    "us-east-2",
				"TableName": "hello",
			}

			_, err := bigtable.Listtables(listtables)

			if err != nil {
				t.Errorf("Test Fail")
			}
		}



    func TestDeletetables(t *testing.T) {

      var dynamodb Dynamodb

      deletetables := map[string]interface{}{
      "Region": "us-east-2",
      "TableName" : "hello",
      }

      _, err := dynamodb.Deletetables(deletetables)

      if err != nil {
        t.Errorf("Test Fail")
      }
    }
