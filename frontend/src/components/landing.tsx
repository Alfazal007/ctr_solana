import { UserContext } from "@/context/UserContext"
import { useContext, useEffect } from "react"
import { useNavigate } from "react-router-dom"
import Navbar from "./Navbar"

const Landing = () => {
	const { user } = useContext(UserContext)
	const navigate = useNavigate()

	useEffect(() => {
		if (!user) {
			navigate("/signin")
			return
		}
	}, [])

	return (
		<>
			{
				!user ?
					<div></div> :
					<>
						<Navbar
							userType={user.userType}
						/>
						<div>Landing</div >
					</>
			}
		</>
	)
}

export default Landing
