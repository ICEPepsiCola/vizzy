import {useEffect, useState} from 'react'
import {useQuery, QueryKey, QueryFunction} from '@tanstack/react-query'

type UsePollingQueryOptions<T> = {
    queryFn: QueryFunction<T>;
    queryKey: QueryKey;
    pollingInterval?: number;
}

function usePollingQuery<T>({queryKey, queryFn, pollingInterval = 10000}: UsePollingQueryOptions<T>) {
    const [pollingTimeoutId, setPollingTimeoutId] = useState<null | number>(null)

    const startPolling = () => {
        const timeoutId = setTimeout(async () => {
            await refetch()
            setPollingTimeoutId(null)
        }, pollingInterval)
        setPollingTimeoutId(timeoutId)
    }

    const stopPolling = () => {
        if (pollingTimeoutId) {
            clearTimeout(pollingTimeoutId)
            setPollingTimeoutId(null)
        }
    }
    useEffect(() => {
        startPolling()
    }, [])
    const {isLoading, error, data, refetch} = useQuery(queryKey, queryFn, {

    })

    return {
        isLoading,
        error,
        data,
        startPolling,
        stopPolling,
    }
}

export default usePollingQuery
