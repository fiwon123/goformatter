package logger

type Logger struct {
	DisableLog bool
}

func (instance *Logger) newLogger(disableLog bool, pretty bool) {

	// instance.DisableLog = disableLog

	// if pretty {
	// 	logging.basicConfig(level=logging.INFO, handlers=[RichHandler()])
	// 	instance.logger = logging.getLogger("rich")
	// } else {
	// 	instance.logger = logging.getLogger("log")
	// }

	// formatter = logging.Formatter(
	// 	"%(asctime)s - %(levelname)s - %(message)s"
	// )

	// instance.logger.setLevel(logging.INFO)

	// if !disadisableLog {
	// 	log_dir = get_path("logs")
	// 	log_dir.mkdir(exist_ok=True)
	// 	timestamp = datetime.datetime.now().strftime("%Y%m%d_%Hh%M")
	// 	file_handler = logging.FileHandler(log_dir / f"formatter_{timestamp}.log")
	// 	file_handler.setFormatter(formatter)
	// 	self.logger.addHandler(file_handler)
	// }
}

func (instance *Logger) saveLog(msg string) {
	// if instance.logger == None || instance.DisableLog {
	// 	return
	// }
	// instance.logger.info(msg)
}

func (instance *Logger) getLogger() {
	// return instance.logger
}
