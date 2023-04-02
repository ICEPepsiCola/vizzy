import {QueryKey, QueryFunction} from '@tanstack/react-query'
import {ComponentType, CSSProperties} from 'react'
import usePollingQuery from '../hook/usePollingQuery'
import {BounceLoader} from 'react-spinners'

type BoxProps = {
    queryKey: QueryKey
    queryFn: QueryFunction
    children: ComponentType<any>
    style?: CSSProperties
}
export default function Box(props: BoxProps) {
    const {queryKey, queryFn, style} = props
    const {isLoading, data} = usePollingQuery({queryKey, queryFn})

    if (isLoading) {
        return <div style={{
            width: style?.width,
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center'
        }}>
            <BounceLoader color='#fff'/>
        </div>
    }
    return (
        <div className='box' style={style}>
            <props.children data={data}/>
        </div>
    )
}
