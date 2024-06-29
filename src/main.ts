import { bootstrapApplication } from '@angular/platform-browser'
import { GlobalCacheConfig, LocalStorageStrategy } from 'ts-cacheable'
import { AppComponent } from './app/app.component'
import { appConfig } from './app/app.config'

// set global cache storage strategy
GlobalCacheConfig.globalCacheKey = 'MM_LOCAL_CACHE'
GlobalCacheConfig.storageStrategy = LocalStorageStrategy
GlobalCacheConfig.maxAge = 600000

setTimeout(() => {
    bootstrapApplication(AppComponent, appConfig).catch(err => {
        console.error(err)
    })
})
