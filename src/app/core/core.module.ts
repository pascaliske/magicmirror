import { NgModule } from '@angular/core'
import { StoreModule } from '@ngrx/store'
import { EffectsModule } from '@ngrx/effects'
import { StoreDevtoolsModule } from '@ngrx/store-devtools'
import { SentryModule } from '@pascaliske/ngx-sentry'
import { NgProgressModule } from 'ngx-progressbar'
import { NgProgressHttpModule } from 'ngx-progressbar/http'
import { NgProgressRouterModule } from 'ngx-progressbar/router'
import { environment } from 'environments/environment'
import { reducers } from 'store'
import { SettingsEffects } from 'store/settings'

@NgModule({
    imports: [
        StoreModule.forRoot(reducers, {
            runtimeChecks: {
                strictActionImmutability: true,
                strictStateImmutability: true,
            },
        }),
        EffectsModule.forRoot([SettingsEffects]),
        StoreDevtoolsModule.instrument({
            name: 'magicmirror',
            logOnly: environment.production,
            autoPause: true,
            maxAge: 25,
        }),
        SentryModule.forRoot({
            enabled: environment.production,
            sentry: environment.sentry,
        }),
        NgProgressModule.withConfig({
            color: '#fff',
            speed: 250,
            thick: true,
            spinner: true,
            meteor: true,
        }),
        NgProgressHttpModule,
        NgProgressRouterModule,
    ],
    exports: [NgProgressModule],
})
export class CoreModule {}
