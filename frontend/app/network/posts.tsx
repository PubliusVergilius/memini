// src/api/axios.ts
import { MutationFunction } from "@tanstack/react-query";
import axios, { AxiosResponse } from "axios";

export type Post = {
	userId: number,
	id: number,
	title: string,
	body: string,
}

const axiosInstance = axios.create({
  baseURL: "https://jsonplaceholder.typicode.com",
  headers: {
    "Content-Type": "application/json",
  },
});

export async function getAllPosts () {
	const response: AxiosResponse<Array<Post>> = await axiosInstance.get("/posts")
	return response.data;
}

export const createPost: MutationFunction<AxiosResponse<Post>> = async (post: Partial<Post>) => {
	const response =  await axiosInstance.post<Post[]>("/posts", post);
	return response;
}
