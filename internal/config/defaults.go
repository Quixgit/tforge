package config

func Default() Config {
	return Config{
		Engine: "auto",

		Terraform: EngineConfig{
			Binary: "terraform",
		},
		Tofu: EngineConfig{
			Binary: "tofu",
		},
		Terragrunt: EngineConfig{
			Binary: "terragrunt",
		},

		UI: UIConfig{
			Theme: "tfui",
		},

		Security: SecurityConfig{
			MaskSecrets:    true,
			AllowApply:     true,
			AllowDestroy:   false,
			PluginMode:     "builtin-only",
			ConfirmApply:   true,
			ConfirmDestroy: true,
		},
	}
}
