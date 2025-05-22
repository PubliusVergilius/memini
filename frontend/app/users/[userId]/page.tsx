import { getUserById } from "@/app/network/users";
import { Params } from "next/dist/server/request/params";
import { notFound } from "next/navigation";

export default async function UserPage({params}: {params: Params}) {
	const { userId } = await params;
	if (!userId) { throw new Error("param for userId is not defined!") };
	const user = await getUserById(Number(userId));

	if (!user) {
		notFound()
	}

	return (
		<main className="py-4">
			<h1 className="text-2xl text-center">User Page</h1>
			<div className="px-8">
				<p className="text-xl"><span className="font-bold">Username</span>: {user.username}</p>
				<p className="text-xl"><span className="font-bold">Email</span>: {user.email}</p>
			</div>
		</main>
	)
}
