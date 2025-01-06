import { UserContext } from "@/context/UserContext"
import { useContext, useEffect } from "react"
import { useNavigate } from "react-router-dom"
import axios from "axios"

const Logout = () => {
	const { user, setUser } = useContext(UserContext)
	const navigate = useNavigate()

	async function logout() {
		await axios.post("http://localhost:8000/api/v1/user/logout", {}, {
			withCredentials: true
		})
		setUser(null)
		navigate("/signin")
	}

	useEffect(() => {
		if (!user) {
			navigate("/signin")
			return
		}
		logout()
	}, [])

	return (
		<div>Logging out</div>
	)
}

export default Logout
