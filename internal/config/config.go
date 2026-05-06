package config

type Config struct {
	Engine string `yaml:"engine"`

	Terraform  EngineConfig `yaml:"terraform"`
	Tofu       EngineConfig `yaml:"tofu"`
	Terragrunt EngineConfig `yaml:"terragrunt"`

	UI       UIConfig       `yaml:"ui"`
	Security SecurityConfig `yaml:"security"`
}

type EngineConfig struct {
	Binary string `yaml:"binary"`
}

type UIConfig struct {
	Theme string `yaml:"theme"`
}

type SecurityConfig struct {
	MaskSecrets    bool   `yaml:"mask_secrets"`
	AllowDestroy   bool   `yaml:"allow_destroy"`
	PluginMode     string `yaml:"plugin_mode"`
	ConfirmApply   bool   `yaml:"confirm_apply"`
	ConfirmDestroy bool   `yaml:"confirm_destroy"`
}
