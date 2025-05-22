import axios, { AxiosResponse } from "axios";

export type User = {
	id: number,
	username: string,
	email: string,
}

const axiosInstance = axios.create({
  baseURL: "https://dummyjson.com",
  headers: {
    "Content-Type": "application/json",
  },
});

export async function getAllUsers () {
	const response: AxiosResponse<Array<User>> = await axiosInstance.get("/users?limit=5")
	return response.data;
}

export async function getUserById (id: number) {
	const response: AxiosResponse<User> = await axiosInstance.get(`/users/${id}`)
	return response.data;
}



