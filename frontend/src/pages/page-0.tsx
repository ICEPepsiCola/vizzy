import Box from '../components/Box'
import request from '../utils/request'
import AnimatedNumbers from 'react-animated-numbers'

const queryFn = () => request({url: '/page-0'})
export default function () {

    return <Box queryKey={['page-0']} queryFn={queryFn} style={{
        width: '100%',
    }}>
        {({data}) => <>
            <div className='card-title'>整体情况</div>
            <div style={{display: 'flex'}}>
                <div style={{width: '25%', textAlign: 'center'}}>
                    <div style={{fontSize: '12px', padding: '10px 0'}}>计算器资源消耗</div>
                    <div style={{
                        display: 'flex',
                        alignItems: 'baseline',
                        justifyContent: 'center',
                    }}>
                        <AnimatedNumbers
                            animateToNumber={data.num1}
                            fontStyle={{fontSize: 32}}
                            configs={(number, index) => {
                                return {mass: 1, tension: 230 * (index + 1), friction: 140}
                            }}
                        ></AnimatedNumbers>
                        <span>GB</span>
                    </div>
                </div>
                <div style={{width: '25%', textAlign: 'center'}}>
                    <div style={{fontSize: '12px', padding: '10px 0'}}>磁盘总空间/利用率</div>
                    <div style={{
                        display: 'flex',
                        alignItems: 'baseline',
                        justifyContent: 'center',
                    }}>
                        <AnimatedNumbers
                            animateToNumber={data.num2}
                            fontStyle={{fontSize: 32}}
                            configs={(number, index) => {
                                return {mass: 1, tension: 230 * (index + 1), friction: 140}
                            }}
                        ></AnimatedNumbers>
                        <span>%</span>
                    </div>
                </div>
                <div style={{width: '25%', textAlign: 'center'}}>
                    <div style={{fontSize: '12px', padding: '10px 0'}}>内存总空间/利用率</div>
                    <div style={{
                        display: 'flex',
                        alignItems: 'baseline',
                        justifyContent: 'center',
                    }}>
                        <AnimatedNumbers
                            animateToNumber={data.num3}
                            fontStyle={{fontSize: 32}}
                            configs={(number, index) => {
                                return {mass: 1, tension: 230 * (index + 1), friction: 140}
                            }}
                        ></AnimatedNumbers>
                        <span>%</span>
                    </div>
                </div>
                <div style={{width: '25%', textAlign: 'center'}}>
                    <div style={{fontSize: '12px', padding: '10px 0'}}>在线节点/利用率</div>
                    <div style={{
                        display: 'flex',
                        alignItems: 'baseline',
                        justifyContent: 'center',
                    }}>
                        <AnimatedNumbers
                            animateToNumber={data.num4}
                            fontStyle={{fontSize: 32}}
                            configs={(number, index) => {
                                return {mass: 1, tension: 230 * (index + 1), friction: 140}
                            }}
                        ></AnimatedNumbers>
                        <span>%</span>
                    </div>
                </div>
            </div>
        </>}
    </Box>
}
