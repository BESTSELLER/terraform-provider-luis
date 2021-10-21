provider "luis" {
  subscription_key = "<yourkey>"
  app_id           = "<yourappid>"
  endpoint         = "westeurope.api.cognitive.microsoft.com" # Defaults to westeurope.api.cognitive.microsoft.com
  luis_version     = "1.0"
}
