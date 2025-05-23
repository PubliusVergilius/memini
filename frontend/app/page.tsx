import React from "react";
import Tasks from "./components/Tasks.tsx"

export default function Notes() {
  return (
	  <div className="border-t border-t-neutral-300 flex flex-row">
		  <div className="border-r border-r-neutral-300 basis-1/3 min-h-screen">
		  	<ul className="flex flex-col items-center px-4">
				<Item text={"Work"} color="bg-yellow-200"/>
				<Item text={"Personal"} color="bg-cyan-200"/>
				<Item text={"College"} color="bg-pink-200"/>
		  	</ul>
		  </div>
		  <Tasks />
		  <div className="basis-1/3"></div>
	  </div>
  );
}

function Item ({ text, color }: { text: string, color: string }) {
	return (<li className="flex-1 min-w-full my-2">
			<div className={`${color} min-w-full px-4 py-4 rounded-xl`} >
				<span className="text-2xl p-4 font-bold">{text}</span>
			</div>
		</li>)
}

