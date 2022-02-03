package config

import "github.com/luthermonson/tutorials-2022-01-cue/config"

#config: config.#Config & {
    app: config.#App & {
        host: #host
        port: 443
        workers: #workers
        dsn: #dsn
    }
}

#config