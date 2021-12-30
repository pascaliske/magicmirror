import { enableProdMode } from '@angular/core'
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic'
import { GlobalCacheConfig, LocalStorageStrategy } from 'ts-cacheable'
import { environment } from 'environments/environment'
import { AppModule } from './app/app.module'

// set global cache storage strategy
GlobalCacheConfig.globalCacheKey = 'MM_LOCAL_CACHE'
GlobalCacheConfig.storageStrategy = LocalStorageStrategy

// enable prod mode
if (environment.production) {
    // eslint-disable-next-line no-console
    console.log(APP_VERSION)
    enableProdMode()
}

platformBrowserDynamic()
    .bootstrapModule(AppModule)
    .catch(err => console.error(err))
