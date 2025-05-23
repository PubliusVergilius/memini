"use client"
import Link from "next/link";
import { usePathname } from "next/navigation";

export default function Header() {
	return (
		<header className="p-4 flex gap-4 justify-between">
			  <Link href={"/"}>
				  <span className="text-3xl font-bold">Notes</span>
			  </Link>
			  <div className="flex gap-8"> 
			    <Link href={"/users"} className={styleLink("/users")}><span className="text-center text-lg">Help</span></Link>
			    <Link href={"/about"} className={styleLink("/about")}><span className="text-center text-lg">About</span></Link>
			  </div> 
		</header>
	)
}

function styleLink(path: string) {
	const _path = usePathname()
	if (path === _path) {
		return "underline";
	}
	return "no-underline"
}
