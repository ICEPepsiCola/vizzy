import ReactECharts from 'echarts-for-react'
import Box from '../components/Box'
import request from '../utils/request'

const queryFn = () => request({url: '/page-3'})
export default function () {
    return <Box queryKey={['page-3']} queryFn={queryFn} style={{
        width: '20%'
    }}>
        {({data}) =>  <>
            <div className='card-title'>平均容量</div>
            <ReactECharts option={{...data, backgroundColor: '#262942'}} theme='dark' style={{
                height: 400,
            }} opts={{
                height: 400,

            }}/>
        </>}
    </Box>
}
