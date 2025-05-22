import { MutationFunction, QueryClientConfig, QueryKey, useMutation, useQuery, useQueryClient } from "@tanstack/react-query";

type UseApiGetType<T> = {
	key: QueryKey
	fn: MutationFunction<T>,
	options?: QueryClientConfig
}

export function useApiGet<T>({key, fn, options}: UseApiGetType<T>) {
	return useQuery({queryKey: key, queryFn: fn, ...options})
}

type UseApiSendType<T> = {
	fn: MutationFunction,
	success?: (data: unknown) => string,
	error?: () => void,
	invalidatedKeys: readonly string[], 
	options?: QueryClientConfig,
}

export function useApiSend<T> ({fn, success, error, invalidatedKeys, options}: UseApiSendType<T>) {
	const queryClient = useQueryClient()
	/*
	const [messageApi, contextHolder] = message.useMessage()
	*/

	const mutation = useMutation({
		mutationFn: fn,		
		onSuccess: data => {
			let message = "Operacao feita com succeso!"
			if (success) {
				message = success(data)
			}

			/*
			messageApi.open({
				type: "success",
				content: message 
			});
			*/

			if (invalidatedKeys.length > 0) { 
				queryClient.invalidateQueries({ queryKey: invalidatedKeys })
			}
		},
		onError: (errorData) => {

			if (error){ error() };
			/*
			messageApi.open({
				type: "error",
				content: "Error!",
			})
			*/
		},
		...options,	
	});

	return {contextHolder, ...mutation};
}

export function getMessageType(httpStatusCode: number): string {
	const prefix = String(httpStatusCode)[0]
	switch (prefix) {
		case "2":
			return "sucess"
		case "4":
			return "error"
		case "5":
			return "error"
		default:
			return "info"
	}
	
}
