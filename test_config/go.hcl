go_module "switchboard" {
  path     = "switchboard"
  upstream = "https://github.com/nint8835/switchboard"
}

go_module "sorik" {
  path     = "sorik"
  upstream = "https://github.com/fogo-sh/sorik"
}

go_module "sorik/borik" {
  path        = "sorik/borik"
  upstream    = "https://github.com/fogo-sh/borik"
  description = "Best discord bot"
}

go_module "testing" {
  path         = "testing"
  upstream     = "https://github.com/nint8835/dieppe-sandbox"
  display_name = "dieppe-sandbox"
  description  = "Testing module for Dieppe"

  readme = <<EOF
Testing module for Dieppe.

# H1

## H2

### H3

#### H4

Ordered list

1. One
2. Two
3. Three
EOF

  link {
    text = "Source"
    url  = ""
  }

  link {
    text = "Docs"
    url  = ""
  }

  link {
    text = "Issues"
    url  = ""
  }

  link {
    text = "Discussions"
    url  = ""
  }

  link {
    text = "Releases"
    url  = ""
  }

  link {
    text = "Actions"
    url  = ""
  }

  link {
    text = "Wiki"
    url  = ""
  }

  link {
    text = "Security"
    url  = ""
  }
}
