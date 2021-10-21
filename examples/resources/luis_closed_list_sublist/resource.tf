resource "luis_closed_list_sublist" "test" {
  closed_list_id = "<listid>"
  canonical_form = "green"
  list           = ["apples", "pears"]
}
