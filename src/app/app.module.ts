import { NgModule, LOCALE_ID } from '@angular/core'
import { CommonModule } from '@angular/common'
import { HttpClientModule } from '@angular/common/http'
import { BrowserModule } from '@angular/platform-browser'
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'
import { ServiceWorkerModule } from '@angular/service-worker'
import { environment } from 'environments/environment'
import { CoreModule } from 'core/core.module'
import { SharedModule } from 'shared/shared.module'
import { AppRoutingModule } from './app-routing.module'
import { AppComponent } from './app.component'

@NgModule({
    imports: [
        CommonModule,
        HttpClientModule,
        BrowserModule,
        BrowserAnimationsModule,
        ServiceWorkerModule.register('ngsw-worker.js', {
            enabled: environment.production,
            registrationStrategy: 'registerWhenStable:30000',
        }),
        CoreModule,
        SharedModule,
        AppRoutingModule,
    ],
    declarations: [AppComponent],
    providers: [
        {
            provide: LOCALE_ID,
            useValue: 'de',
        },
    ],
    bootstrap: [AppComponent],
})
export class AppModule {}
