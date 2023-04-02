import Box from '../components/Box'
import request from '../utils/request'
import Table from 'rc-table'
import {ColumnsType} from 'rc-table/es/interface'

const queryFn = () => request({url: '/page-5'})
export default function () {
    const columns: ColumnsType<any> = [
        {
            title: '地区',
            dataIndex: 'area',
            key: 'area',
            align: 'center',
        },
        {
            title: '统计',
            dataIndex: 'count',
            key: 'count',
            align: 'center',
        },
        {
            title: '百分比',
            dataIndex: 'percent',
            key: 'percent',
            align: 'center',
        },
        {
            title: '时长',
            dataIndex: 'duration',
            key: 'duration',
            align: 'center',
        },
    ]
    return <Box queryKey={['page-5']} queryFn={queryFn} style={{
        width: '25%'
    }}>
        {({data}) => <>
            <div className='card-title'>运行任务</div>
            <Table columns={columns} data={data} rowKey='id' scroll={{
                x: 'max-content',
            }}/>
        </>}
    </Box>
}
