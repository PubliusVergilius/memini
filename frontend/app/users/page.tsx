import { Suspense } from "react"
import { getAllUsers } from "../network/users";
import LoadingUsers from "./loading-old";
import Link from "next/link";

async function Users() {
	const data = await getAllUsers();

	/**
	* remove the comment below to see the error page with a message.
	*/
	// throw new Error("error xyz")
	return (<ul>
			{data.users.map(user =>
				<li key={user.id}>
					<div>
						<Link href={`/users/${user.id}`}>
							{user.username}
						</Link>
					</div>
				</li>
		       )}
	       </ul>)
}

export default async function UsersPage () {
	return (
		<main>
			<Suspense fallback={<LoadingUsers />}>
				<Users />
			</ Suspense>
		</main>
	)
}
