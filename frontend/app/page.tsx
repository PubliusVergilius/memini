import React from "react";

export default function Notes() {
  return (
	  <div className="border-t border-t-neutral-300 flex flex-row">
		  <div className="border-r border-r-neutral-300 basis-1/3 min-h-screen">
		  	<ul className="flex flex-col items-center">
				<Item text={"Work"} color="bg-yellow-200"/>
				<Item text={"Personal"} color="bg-cyan-200"/>
				<Item text={"College"} color="bg-pink-200"/>
		  	</ul>
		  </div>
		  <div className="basis-1/3"></div>
	  </div>
  );
}

function Item ({ text, color }: { text: string, color: string }) {
	return <li className="flex-1"><div className={`${color} min-w-full px-4 py-2 rounded-md`}  >{text}</div></li>
}

