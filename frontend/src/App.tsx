import darkTheme from './assets/dark.json'
import * as echarts from 'echarts'
import {QueryClient, QueryClientProvider} from '@tanstack/react-query'
// import RGL, {WidthProvider} from 'react-grid-layout'
import Page0 from './pages/page-0'
import Page1 from './pages/page-1'
import Page2 from './pages/page-2'
import Page3 from './pages/page-3'
import Page4 from './pages/page-4'
import Page5 from './pages/page-5'
import Page6 from './pages/page-6'
import './App.less'
import {useEffect} from 'react'

echarts.registerTheme('dark', darkTheme)
const queryClient = new QueryClient()

// const ReactGridLayout = WidthProvider(RGL)

function App() {
    useEffect(() => {
        // @ts-ignore
        document.documentElement.style.zoom = `${window.innerWidth / 1920}`
    })
    return <div className='app'>
        <div className='app-title'>数据计算服务屏</div>
        {/*<ReactGridLayout>*/}
        <div className='app-column' style={{
            width: '100%',
            height: '120px',
            display: 'flex',
        }}>
            <Page0/>
        </div>
        <div className='app-column' style={{
            width: '100%',
            height: '440px',
            display: 'flex',
            paddingBottom: '10px',
        }}>
            <Page1/>
            <Page2/>
            <Page3/>
        </div>
        <div className='app-column' style={{
            width: '100%',
            height: '440px',
            display: 'flex',
        }}>
            <Page4/>
            <Page5/>
            <Page6/>
        </div>
        {/*</ReactGridLayout>*/}
    </div>
}

export default function () {
    return <QueryClientProvider client={queryClient}>
        <App/>
    </QueryClientProvider>
}
