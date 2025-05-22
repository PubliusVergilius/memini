"use client"

export default function UserPageError({error}: {error: Error}) {
	return (<main className="py-4">
			<h2 className="text-2xl text-center">An error ocurred!</h2>
			<h2 className="text-xl text-center">Error on loading users data: {error.message}</h2>
		</main>)
}
