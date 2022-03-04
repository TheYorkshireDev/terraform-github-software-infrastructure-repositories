locals {
  stuff   = ["a", "b", "c"]
  instuff = index(local.stuff, "d")
}
