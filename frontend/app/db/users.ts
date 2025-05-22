import sql from "better-sqlite3"

const db = sql("./users.db")
export async function getAllUsers() {
	await new Promise((res) => setTimeout(res, 2000))
	return db.prepare("SELECT * FROM users").all();
}
