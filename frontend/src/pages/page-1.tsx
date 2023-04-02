import ReactECharts from 'echarts-for-react'
import Box from '../components/Box'
import request from '../utils/request'

const queryFn = () => request({url: '/page-1'})
export default function () {
    return <Box queryKey={['page-1']} queryFn={queryFn} style={{
        width: '50%',
    }}>
        {({data}) => {
            console.log(data)
            return <>
                <div className='card-title'>MaxComputed总数</div>
                <ReactECharts option={{...data, backgroundColor: '#262942'}} theme='dark' style={{
                    height: 400,
                }} opts={{
                    height: 400,

                }}/>
            </>
        }}
    </Box>
}
