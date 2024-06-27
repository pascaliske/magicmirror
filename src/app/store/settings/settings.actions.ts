import { createActionGroup, props, emptyProps } from '@ngrx/store'
import { Settings } from './settings.model'

export const SettingsActions = createActionGroup({
    source: 'Settings',
    events: {
        // load
        Load: emptyProps(),
        LoadSuccess: props<{ payload: Settings }>(),
        LoadError: props<{ payload: string; error?: Error }>(),
    },
})
