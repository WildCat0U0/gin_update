package config

type Database struct {
	Driver   string `mapstructure:"driver" json:"driver" yaml:"driver"` // mysql
	Host     string `mapstructure:"host" json:"host" yaml:"host"`       //
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	UserName string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Charset  string `mapstructure:"charset" json:"charset" yaml:"charset"` // utf8mb4
	//Collation string `mapstructure:"collation" json:"collation" yaml:"collation"`
	MaxIdleConns        int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`                         // 最大空闲连接数
	MaxOpenConns        int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`                         // 最大连接数
	LogMode             string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`                                           // 是否开启日志模式
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"` // 是否开启文件日志模式
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`                               // 日志文件名
}
