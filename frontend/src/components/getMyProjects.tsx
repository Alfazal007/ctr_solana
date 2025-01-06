import { Card, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { CheckCircle, PlayCircle, XCircle } from 'lucide-react'
import axios from "axios"
import { toast } from "@/hooks/use-toast"
import { useNavigate } from "react-router-dom"
import { useContext, useEffect, useState } from "react"
import { UserContext } from "@/context/UserContext"

interface Task {
	id: string
	name: string
	started: boolean
	completed: boolean
}

export default function TaskCardList() {
	const navigate = useNavigate()
	const { user } = useContext(UserContext)
	const [tasks, setProjects] = useState<Task[]>([])

	async function getProjects() {
		try {
			const projectsData = await axios.get("http://localhost:8000/api/v1/project/projects", {
				withCredentials: true
			})
			console.log({ projectsData })
			if (projectsData.status != 200) {
				toast({ title: "Issue fetching the data", variant: "destructive" })
				navigate("/")
				return
			}
			if (projectsData.data.length == 0) {
				toast({ title: "Issue fetching the data" })
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
		<div className="min-h-screen bg-gray-900 p-8">
			<div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
				{tasks.map((task) => (
					<Card key={task.id} className="bg-gray-800 text-white">
						<CardContent className="p-6">
							<h3 className="text-xl font-semibold mb-4">{task.name}</h3>
							<div className="flex items-center space-x-2">
								<Badge
									variant={task.started ? "default" : "secondary"}
									className="flex items-center space-x-1"
								>
									{task.started ? (
										<PlayCircle className="w-4 h-4" />
									) : (
										<XCircle className="w-4 h-4" />
									)}
									<span>{task.started ? "Started" : "Not Started"}</span>
								</Badge>
								<Badge
									variant={task.completed ? "default" : "secondary"}
									className="flex items-center space-x-1"
								>
									{task.completed ? (
										<CheckCircle className="w-4 h-4" />
									) : (
										<XCircle className="w-4 h-4" />
									)}
									<span>{task.completed ? "Completed" : "Not Completed"}</span>
								</Badge>
							</div>
						</CardContent>
					</Card>
				))}
			</div>
		</div>
	)
}

