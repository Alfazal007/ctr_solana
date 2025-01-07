import { Card, CardContent } from "@/components/ui/card"
import axios from "axios"
import { toast } from "@/hooks/use-toast"
import { useNavigate } from "react-router-dom"
import { useContext, useEffect, useState } from "react"
import { UserContext } from "@/context/UserContext"
import Navbar from "./Navbar"

interface Task {
	id: string
	name: string
}

export default function TaskCardListToVote() {
	const navigate = useNavigate()
	const { user } = useContext(UserContext)
	const [tasks, setProjects] = useState<Task[]>([])

	async function getProjects() {
		try {
			const projectsData = await axios.get("http://localhost:8000/api/v1/project/projectsToVote", {
				withCredentials: true
			})
			console.log({ projectsData })
			if (projectsData.status != 200) {
				toast({ title: "Issue fetching the data", variant: "destructive" })
				navigate("/")
				return
			}
			if (projectsData.data.length == 0) {
				toast({ title: "Nothing to display" })
				return
			}
			setProjects(projectsData.data)
		} catch (err) {
			toast({ title: "Issue fetching the data" })
			navigate("/")
			return
		}
	}

	useEffect(() => {
		if (!user) {
			navigate("/signin")
			return
		}
		getProjects()
	}, [])

	return (
		<>
			{
				user &&
				<>
					<Navbar userType={user.userType} />
					<div className="min-h-screen bg-gray-900 p-8">
						<div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
							{tasks.map((task) => (
								<Card onClick={() => {
									navigate(`/project/vote/${task.id}`)
								}} key={task.id} className="bg-gray-800 text-white cursor-pointer">
									<CardContent className="p-6">
										<h3 className="text-xl font-semibold mb-4">{task.name}</h3>
									</CardContent>
								</Card>
							))}
						</div>
					</div>
				</>
			}
		</>
	)
}
