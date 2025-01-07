import { toast } from "@/hooks/use-toast"
import axios from "axios"
import { useContext, useEffect, useState } from "react"
import { useNavigate, useParams } from "react-router-dom"
import { Project } from "./addImageToProject"
import { UserContext } from "@/context/UserContext"
import { Button } from "./ui/button"

const EndProject = () => {
	const [project, setProject] = useState<Project | null>(null)
	const { projectId } = useParams();
	const navigate = useNavigate()
	const { user } = useContext(UserContext)

	async function fetchProject() {
		const response = await axios.get(`http://localhost:8000/api/v1/project/${projectId}`, { withCredentials: true })
		if (response.status != 200) {
			toast({ title: "Issue  fetching project" })
			navigate("/")
			return
		}
		setProject(response.data)
		if (project?.completed == true) {
			toast({ title: "The project is already complete" })
			navigate("/")
		}
		if (project?.started == false) {
			toast({ title: "The project has not started yet" })
			navigate("/")
		}
	}

	useEffect(() => {
		if (!user) {
			navigate("/signin")
			return
		}
		fetchProject()
	}, [])

	async function endProject() {
		if (!project || project.completed == true) {
			return
		}
		try {
			const endResponse = await axios.put(`http://localhost:8000/api/v1/project/end/${project.id}`, {}, { withCredentials: true })
			if (endResponse.status != 200) {
				toast({ title: "Issue ending the project", variant: "destructive" })
				return
			}
			toast({ title: "Ended the project successfully" })
			navigate("/")
		} catch (err) {
			toast({ title: "Issue ending the project", variant: "destructive" })
		}
	}

	return (
		<div className="min-h-screen flex items-center justify-center bg-gradient-to-b from-gray-900 to-gray-800">
			<Button
				onClick={endProject}
				size="lg"
				className="bg-blue-900 text-white border-none hover:bg-blue-900 hover:text-white"
			>
				End Project
			</Button>
		</div>
	)
}

export default EndProject
