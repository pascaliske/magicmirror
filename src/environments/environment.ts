import 'zone.js/plugins/zone-error'

export const environment = {
    production: false,
    version: APP_VERSION,
    sentry: {
        dsn: 'https://47059a8f76b54870afbd05bf26d72266@sentry.io/1333718',
        environment: 'development',
        release: APP_VERSION,
    },
}
