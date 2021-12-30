import { InjectionToken } from '@angular/core'
import { parse, format } from 'date-fns'
import { de } from 'date-fns/locale'

// joined type of parse() options and format() options
export type DateFnsConfig = Parameters<typeof parse>[3] & Parameters<typeof format>[2]

// injection token for global date-fns config
export const DATE_FNS_CONFIG = new InjectionToken<DateFnsConfig>('date', {
    providedIn: 'root',
    factory: () => ({
        locale: de,
    }),
})
