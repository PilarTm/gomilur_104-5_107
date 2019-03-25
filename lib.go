package milur104

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"github.com/martinlindhe/crc24"
	"strings"
)

const (
	CODE_VOLTAGE                       = 0x01 // Напряжение
	CODE_CURRENT                       = 0x02 // Ток
	CODE_ACTIVE_POWER                  = 0x03 // Активная мощность
	CODE_ACTIVE_ENERGY_TARIFF_SUM      = 0x04 // Активная энергия суммарная
	CODE_ACTIVE_ENERGY_TARIFF_1        = 0x05 // Активная энергия по тарифу 1
	CODE_ACTIVE_ENERGY_TARIFF_2        = 0x06 // Активная энергия по тарифу 2
	CODE_ACTIVE_ENERGY_TARIFF_3        = 0x07 // Активная энергия по тарифу 3
	CODE_ACTIVE_ENERGY_TARIFF_4        = 0x08 // Активная энергия по тарифу 4
	CODE_FREQUENCY                     = 0x09 // Частота сети
	CODE_CHECK_TARIFF                  = 0x0A // Текущий тариф
	CODE_PARAMERTERS_IDENT             = 0x0B // Параметры индикации
	CODE_IDENT_MANAGE_PROCEDURE        = 0x0D // Идентификатор управляющей процедуры
	CODE_CLOCK_REAL_TIME               = 0x0E // Часы реального времени
	CODE_HOLIDAY_LIST                  = 0x0F // Список праздничных дней
	CODE_SLICE_POWER                   = 0x10 // Срезы мощности
	CODE_BUF_ACTIONS                   = 0x11 // Буфер событий
	CODE_BUF_ERRORS                    = 0x12 // Список событий (errors)
	CODE_MAX_COUNT_TARIFFS             = 0x13 // Максимальное число тарифов
	CODE_TARIFFS_CRON_JAN              = 0x14 // Тарифное расписание на январь
	CODE_TARIFFS_CRON_FEB              = 0x15 // Тарифное расписание на февраль
	CODE_TARIFFS_CRON_MAR              = 0x16 // Тарифное расписание на март
	CODE_TARIFFS_CRON_APR              = 0x17 // Тарифное расписание на апрель
	CODE_TARIFFS_CRON_MAY              = 0x18 // Тарифное расписание на май
	CODE_TARIFFS_CRON_JIN              = 0x19 // Тарифное расписание на июнь
	CODE_TARIFFS_CRON_JIL              = 0x1A // Тарифное расписание на июль
	CODE_TARIFFS_CRON_AUG              = 0x1B // Тарифное расписание на август
	CODE_TARIFFS_CRON_SEN              = 0x1C // Тарифное расписание на сентябрь
	CODE_TARIFFS_CRON_OKT              = 0x1D // Тарифное расписание на октябрь
	CODE_TARIFFS_CRON_NOV              = 0x1E // Тарифное расписание на ноябрь
	CODE_TARIFFS_CRON_DEC              = 0x1F // Тарифное расписание на декабрь
	CODE_MORE_INFO                     = 0x20 // Информация об устройстве (наименование, производитель)
	CODE_VERSION_SOFT                  = 0x21 // Версия встроенного программного обеспечения
	CODE_COLIBRATION_REAL_TIME         = 0x22 // Калибровка часов реального времени
	CODE_CONTROL_WINTER_SUMER          = 0x23 // Управление автоматическим переходом на летнее/зимнее время
	CODE_LOWER_LIMIT_POWER             = 0x24 // Нижний предел по напряжению
	CODE_UPPER_LIMIT_POWER             = 0x25 // Верхний предел по напряжению
	CODE_LOWER_LIMIT_FREQUENCY         = 0x26 // Нижний предел по частоте
	CODE_UPPER_LIMIT_FREQUENCY         = 0x27 // Верхний предел по частоте
	CODE_LOWER_LIMIT_ACTIVE_POWER      = 0x28 // Верхний предел по активной мощности
	CODE_SETTIONS_CONNECTION           = 0x29 // Параметры сеанса связи
	CODE_PASSWORD_LVL1                 = 0x2A // Пароль 1-го уровня
	CODE_PASSWORD_LVL2                 = 0x2B // Пароль 2-го уровня
	CODE_PASSWORD_LVL3                 = 0x2C // Пароль 3-го уровня
	CODE_CRON_CONTROL_00_06_MONTH      = 0x36 // Расписание управления освещением на первое полугодие
	CODE_CRON_CONTROL_06_12_MONTH      = 0x37 // Расписание управления освещением на первое полугодие
	CODE_SECURE_OFF_POWER              = 0x38 // Защитное отключение нагрузки
	CODE_VOLTAGE_BATTARY               = 0x39 // Напряжение батареи резервного питания
	CODE_TECHNICAL_OBJECT              = 0x3A // Технологический объект
	CODE_LIST_MESSAGED                 = 0x3B // Список событий (messages)
	CODE_LIST_WARNING                  = 0x3C // Список событий (warnings)
	CODE_TIME_INTERGRATION_POWER_POROF = 0x3D // Время интегрирования профиля мощности
	CODE_NUMBER_SOFTWARE               = 0x3E // Цифровой идентификатор ПО
	CODE_MODE_IMP_OUT                  = 0x3F // Режим импульсного выхода счётчика
	CODE_TYPE_CONTROL_POWER            = 0x40 // Тип выхода управления нагрузкой
	CODE_LIMIT_AUTO_OFF_POWER          = 0x41 // Порог автоматического отключения нагрузки
	CODE_ENERGY_DAILY                  = 0x42 // Энергия в суточных интервалах
	CODE_ENERGY_MONTHLY                = 0x43 // Энергия в месячных интервалах
	CODE_SERIAL_NUMBER                 = 0x44 // Серийный номер счетчика
	CODE_TYPE_ADDRESSES                = 0x55 // Тип адресации
	CODE_KEY_ZIGBEE                    = 0x56 // Ключ сети ZigBee
	CODE_CONF_DEVICE                   = 0x57 // Конфигурация счетчика
	// Ниже: Объект используется в ПО счетчика, начиная с версии 1.12 включительно.
	CODE_PARAMERTERS_COLIBRATION_bef_3_0 = 0x0C // Параметры калибровки. Объект используется в ПО счетчика
	CODE_TIMEOUT_RESPONSE                = 0x45 // Таймаут ответа счетчика
	CODE_VERSION_METROLOGY_SOFTWARE      = 0x53 // Версия метрологически значимой части ПО
	// Ниже: Объект используется в ПО счетчика, начиная с версии 3.00 включительно.
	CODE_PARAMERTERS_COLIBRATION_aft_3_0 = 0x46 // Параметры калибровки
	CODE_SERIAL_PRIINT_PONT              = 0x47 // Серийный номер печатного узла
	CODE_MORE_PARAMETERS                 = 0x4F // Дополнительные параметры.
	// Ниже: Объект используется в ПО счетчика, начиная с версии 4.00 включительно.
	CODE_REACTIVE_ENERGY_SUM        = 0x48 // Реактивная энергия суммарная
	CODE_REACTIVE_ENERGY_TARIFF_1   = 0x49 // Реактивная энергия по тарифу 1
	CODE_REACTIVE_ENERGY_TARIFF_2   = 0x4A // Реактивная энергия по тарифу 2
	CODE_REACTIVE_ENERGY_TARIFF_3   = 0x4B // Реактивная энергия по тарифу 3
	CODE_REACTIVE_ENERGY_TARIFF_4   = 0x4C // Реактивная энергия по тарифу 4
	CODE_REACTIVE_POWER             = 0x4D // Реактивная мощность
	CODE_FULL_POWER                 = 0x4E // Полная мощность
	CODE_FACTOR_POWER               = 0x50 // Фактор мощности
	CODE_PARAMERTERS_COLIBRATION_I3 = 0x51 // Параметры калибровки канала I3
	CODE_IDENT_RUN_CURRENT_CHANEL   = 0x52 // Идентификатор текущего токового канала
	CODE_CONTROL_RELE_POWER         = 0x54 // Управление встроенным реле отключения нагрузки
)

const (
	SET_RESET_SETTINGS                = 0x01 // Сброс на заводские установки
	SET_RESET_BUF_ERROR               = 0x02 // Сброс буфера ошибок
	SET_RESET_ENERGY                  = 0x03 // Сброс накопленной энергии. Должна быть установлена защитная перемычка.
	SET_REBOOT_CALIBRATIONCOEFS       = 0x04 // Перезагрузка калибровочных коэффициентов в измерительный модуль
	SET_APPLY_MODE_SETTINS_CONNECTION = 0x05 // Применить изменения в интерфейсном объекте "Параметры сеанса связи".
	SET_RUN_CHANGESBATTARY_VOLTAGE    = 0x06 // Однократный запуск измерения напряжения батареи резервного питания
	SET_MANUF_INIT_DEVICE             = 0x07 // Заводская инициализация счетчика. Должна быть установлена защитная перемычка.
	SET_VIEW_ADDR_DEVICE              = 0x08 // Показать адрес счетчика
)

const (
	ERROR_NO                   = 0x00
	ERROR_ILLEGAL_FUNCTION     = 0x01 // Некорректный идентификатор функции
	ERROR_ILLEGAL_DATA_ADDRESS = 0x02 // Некорректный идентификатор объекта
	ERROR_ILLEGAL_DATA_VALUE   = 0x03 // Некорректное значение данных
	ERROR_SLAVE_DEVICE_FAILURE = 0x04 // Невозможно выполнить команду
	ERROR_ACKNOWLEDGE          = 0x05 // Запрос принят, начата обработка
	ERROR_SLAVE_DEVICE_BUSY    = 0x06 // Устройство занято
	ERROR_EEPROM_ACCESS_ERROR  = 0x07 // Ошибка доступа к памяти EEPROM
	ERROR_SESSION_CLOSED       = 0x08 // Сеанс связи закрыт
	ERROR_ACCESS_DENIED        = 0x09 // Доступ с указанным уровнем запрещён
	ERROR_ERROR_CRC            = 0x0A // Ошибка контрольной суммы
	ERROR_FRAME_INCORRECT      = 0x0B // Некорректный фрейм
	ERROR_JUMPER_ABSENT        = 0x0C // Не установлена защитная перемычка
	ERROR_PASSW_INCORRECT      = 0x0D // Неверный пароль
)

const (
	GRCODE_USER    = 0x00 // уровень пользователя
	GRCODE_ADMIN   = 0x01 // уровень администратора
	GRCODE_DEVELOP = 0x02 // уровень разработчика
)

const (
	SERVICE_AOPEN    = 0x08 // Код сервиса AOPEN
	SERVICE_ARELEASE = 0x09 // Код сервиса ARELEASE
)

const (
	ERROR_STR_INVALID_ACCESS_CODE = "invalid code access"
)

const (
	DEFAULT_PASS = "FFFFFFFFFFFF" // 2*6
)

type Lib struct {
	addr     int
	password string
	version  int
}

func New(add int, pass string) *Lib {
	return &Lib{addr: addr, pass: pass}

}

func Empty() *Lib {
	return &Lib{}

}

func (l *Lib) SetAddr(addr int) {
	l.addr = addr
}

func (l *Lib) SetPassword(pass string) {
	l.pass = pass
}

// Установить значение объекта
func (l *Lib) Set(prop byte, value []byte) []byte {
	res := []byte{}
	return res
}

// получить значение объекта
func (l *Lib) Get(prop byte) []byte {
	res := []byte{}

	return res
}

// Получить байт объекта
func (l *Lib) GetByte(prop byte, index int) []byte {
	res := []byte{}

	return res
}

// Установить байт
func (l *Lib) SetByte(prop byte, index int, value byte) []byte {
	res := []byte{}

	return res
}

// Инициализация списка
func (l *Lib) ListInit(prop byte) []byte {
	res := []byte{}

	return res
}

// Получить число элементов в списке
func (l *Lib) GetListNE(prop byte) []byte {
	res := []byte{}

	return res
}

// получить элемент списка профиля мощности
func (l *Lib) GetListRecPWI(prop byte, index int) []byte {
	res := []byte{}

	return res
}

// открытие сеанса связи. Вызывает сервис AOPEN.
// По умолчанию открывает от администратора с паролем 000000
func (l *Lib) Aopen(gr_code byte, pass string) ([]byte, error) {
	res := []byte{}

	addr := get_addr(l.addr)

	for i := 0; i < len(addr); i++ {
		res = append(res, addr[i])
	}

	if !is_valid_gr_code(gr_code) {
		return []byte{}, errors.New(ERROR_STR_INVALID_ACCESS_CODE)
	}

	if pass == "" {
		pass = DEFAULT_PASS
	}

	pass_to_hex := hex.DecodeString(strings.Join(strings.Fields(pass), ""))
	for i := 0; i < len(pass_to_hex); i++ {
		res = append(res, pass_to_hex[i])
	}

	crc := gen_crc(res)

	for i := 0; i < len(gen_crc); i++ {
		res = append(res, gen_crc[i])
	}

	return res, nil
}

// акрытие сеанса связи. Вызывает сервис ARELEASE
func (l *Lib) Arelease() []byte {
	res := []byte{}

	return res
}

// Установка времени
func (l *Lib) SetRtc(prop, sec, min, hour, day, mon, year byte) []byte {
	res := []byte{}

	return res
}

// получить элемент списка событий
func (l *Lib) GetEvtList(prop byte, index int) []byte {
	res := []byte{}

	return res
}

// получить элемент списка энергий в интервалах
func (l *Lib) GetEntaList(prop byte, index int) []byte {
	res := []byte{}

	return res
}

// получить элемент списка профиля мощности с параметром неполного интервала.
// Данная команда действительна, начиная с версии 1.10 включительно
func (l *Lib) GetListRecPWI_Par(prop byte, index int) []byte {
	res := []byte{}

	return res
}

// найти запись в списке профиля мощности по заданным дате и времени записи.
// Данная команда действительна, начиная с версии 3.00 включительно.
func (l *Lib) PwiListSearch(prop, sec, min, hour, day, mon, year byte) []byte {
	res := []byte{}

	return res
}

// получить конфигурацию и значения адресов счетчика.
// Данный метод доступен, начиная с версии 4.00 включительно
func (l *Lib) GetAddrConfig() []byte {
	res := []byte{}

	return res
}

// найти запись в списке по заданным начальным и конечным дате и времени.
// Данный метод доступен, начиная с версии 4.05 включительно.
func (l *Lib) ListSerach(prop,
	start_min, start_hour, start_day, start_mon, start_year,
	end_min, end_hour, end_day, end_mon, end_year byte) []byte {
	res := []byte{}

	return res
}

// получить коллекцию данных. Метод возвращает данные из нескольких объектов.
// Данный метод доступен, начиная с версии 4.05 включительно.
func (l *Lib) GetCollection(collect_uid byte, index int) []byte {
	res := []byte{}

	return res
}

// SYSTEM FUNC
// Валидный ли уровень доступа
func is_valid_gr_code(gr_code byte) bool {
	return gr_code != GRCODE_USER &&
		gr_code != GRCODE_ADMIN &&
		gr_code != GRCODE_DEVELOP
}

// Подсчет контрольной суммы
func gen_crc(pack []byte) []byte {
	return crc24.Sum(pack)
}

// перобразование адреса
func get_addr(addr int) []byte {
	if addr < 0xFF {
		return []byte{byte(addr)}
	} else {
		addr_c := make([]byte, 4)
		binary.LittleEndian.PutUint32(addr, uint32(l.addr))
		return addr_c
	}
}
