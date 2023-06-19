import { enableProdMode } from '@angular/core'
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic'
import { GlobalCacheConfig, LocalStorageStrategy } from 'ts-cacheable'
import { environment } from 'environments/environment'
import { AppModule } from './app/app.module'

// set global cache storage strategy
GlobalCacheConfig.globalCacheKey = 'MM_LOCAL_CACHE'
GlobalCacheConfig.storageStrategy = LocalStorageStrategy
GlobalCacheConfig.maxAge = 600000

// enable prod mode
if (environment.production) {
    enableProdMode()
}

platformBrowserDynamic()
    .bootstrapModule(AppModule)
    .catch(err => console.error(err))
