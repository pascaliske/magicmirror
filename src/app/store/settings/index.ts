import { createFeature } from '@ngrx/store'
import { settingsReducer } from './settings.reducer'

export * from './settings.actions'
export * from './settings.effects'
export * from './settings.reducer'
export * from './settings.model'

export const SettingsFeature = createFeature({
    name: 'settings',
    reducer: settingsReducer,
    // extraSelectors: ({}) => ({}),
})
