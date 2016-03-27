package utils

import(
    log "linq/core/log"
)

func HandleWarn(err error) {
    if err != nil {
        log.Warn(err.Error())
    }
}

func HandleFatal(err error) {
    if err != nil {
        log.Fatal(err.Error())
    }
}
