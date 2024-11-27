import type { ApplicationConfig, ValueProvider } from '@angular/core'
import {
    provideExperimentalZonelessChangeDetection,
    importProvidersFrom,
    APP_ID,
    LOCALE_ID,
} from '@angular/core'
import { provideRouter } from '@angular/router'
import { provideHttpClient, HttpClient, withFetch } from '@angular/common/http'
import { provideStore } from '@ngrx/store'
import { provideStoreDevtools } from '@ngrx/store-devtools'
import { provideRouterStore } from '@ngrx/router-store'
import { provideEffects } from '@ngrx/effects'
import { SentryModule } from '@pascaliske/ngx-sentry'
import { TranslateModule, TranslateLoader } from '@ngx-translate/core'
import { provideNgProgressOptions } from 'ngx-progressbar'
import { provideNgProgressHttp } from 'ngx-progressbar/http'
import { provideNgProgressRouter } from 'ngx-progressbar/router'
import { environment } from 'environments/environment'
import { reducers, metaReducers, storeDevtoolsOptions } from 'store'
import { SettingsEffects } from 'store/settings'
import { TranslationLoader } from 'core/translation.loader'
import { features, routes } from './app.routing'

export const provideAppId: () => ValueProvider = (): ValueProvider => ({
    provide: APP_ID,
    useValue: 'magicmirror',
})

export const provideLocaleId: () => ValueProvider = (): ValueProvider => ({
    provide: LOCALE_ID,
    useValue: 'de',
})

export const appConfig: ApplicationConfig = {
    providers: [
        provideExperimentalZonelessChangeDetection(),
        importProvidersFrom(
            SentryModule.forRoot({
                enabled: environment.production,
                sentry: environment.sentry,
            }),
            TranslateModule.forRoot({
                defaultLanguage: 'de',
                loader: {
                    provide: TranslateLoader,
                    useClass: TranslationLoader,
                    deps: [HttpClient],
                },
            }),
        ),
        provideRouter(routes, ...features),
        provideHttpClient(withFetch()),
        provideStore(reducers, {
            metaReducers,
            runtimeChecks: {
                strictActionImmutability: true,
                strictStateImmutability: true,
            },
        }),
        provideStoreDevtools(storeDevtoolsOptions),
        provideRouterStore(),
        provideEffects(SettingsEffects),
        provideNgProgressOptions({ speed: 250, spinner: true, flat: true }),
        provideNgProgressHttp({}),
        provideNgProgressRouter({ minDuration: 1500 }),
        provideAppId(),
        provideLocaleId(),
    ],
}
