import ReactECharts from 'echarts-for-react'
import Box from '../components/Box'
import request from '../utils/request'

const queryFn = () => request({url: '/page-4'})
export default function () {
    return <Box queryKey={['page-4']} queryFn={queryFn} style={{
        width: '25%',
    }}>
        {({data}) => <>
            <div className='card-title'>储存容量</div>
            <ReactECharts option={{...data, backgroundColor: '#262942'}} theme='dark' style={{
                height: 400,
            }} opts={{
                height: 400,

            }}/>
        </>}
    </Box>
}
