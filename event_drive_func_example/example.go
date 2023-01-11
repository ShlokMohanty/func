package mycloudeventfunction 
import(
"context"
"github.com/GoogleCloudPlatform/functions-framework-go/functions"
"github.com/cloudevents/sdk-go/v2/event"
)
func init() {
  //register a cloudevent function with the function framework 
  functions.CloudEvent("MyCloudEventFunction", myCloudEventFunction)
}
