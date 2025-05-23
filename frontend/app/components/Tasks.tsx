export default function Tasks () {

	const pallete = {
		yellow: "bg-yellow-200 ",
		blue: "bg-cyan-200 ",
		pink: "bg-pink-200 ",
	};

	const bodyMock = "As armas e os barões assinalados, que da ocidental praia Lusitana, por mares nunca de antes navegados, passaram ainda além de Taprobana.";

	const randomColor = () => {
		const random = Math.floor(Math.random() * 3);
		return Object.values(pallete)[random]
	};

	return (
		<Grid>
			<Card title="Teste 1" body={bodyMock} color={randomColor()}/>
			<Card title="Teste 2" body={bodyMock} color={randomColor()}/>
			<Card title="Teste 3" body={bodyMock} color={randomColor()}/>
			<Card title="Teste 4" body={bodyMock} color={randomColor()}/>
		</Grid>
       );
}

const Grid: React.FC<GridProps> = ({ children }) => {
  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 p-4">
	    {children} 
    </div>
  );
};

function Card ({title, body, color}: {title: string, body: string, color: string}) {

	return (
		<div className={`flex flex-col p-4 rounded-xl ${color}`}>
			<h3 className="text-2xl font-bold">{title}</h3>
			<p className="text-lg">{body}</p>
		</div>
	); 
}
