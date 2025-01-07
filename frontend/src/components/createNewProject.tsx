import { UserContext } from "@/context/UserContext"
import { FormEvent, useContext, useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"
import { Input } from "./ui/input"
import { Button } from "./ui/button"
import axios from "axios"
import { toast } from "@/hooks/use-toast"
import Navbar from "./Navbar"

const CreateNewProject = () => {
	const { user } = useContext(UserContext)
	const navigate = useNavigate()
	const [projectName, setProjectName] = useState("")

	useEffect(() => {
		if (!user) {
			navigate("/signin")
			return
		} else if (user.userType != "creator") {
			navigate("/")
			return
		}
	}, [])

	async function handleSubmit(e: FormEvent) {
		e.preventDefault()
		if (projectName.trim().length > 0 && projectName.trim().length <= 20) {
			const response = await axios.post("http://localhost:8000/api/v1/project/create-project", {
				name: projectName.trim()
			}, {
				withCredentials: true
			})
			console.log({ response })
			if (response.status != 200) {
				toast({
					title: "Issue creating new project",
					description: response.data
				})
				return
			} else {
				navigate("/add-image/" + response.data.id)
				return
			}
		} else {
			toast({
				title: "Project name should be between 1 and 20"
			})
		}
	}

	return (
		<>
			{
				user &&
				<>
					<Navbar userType={user.userType} />
					<div className="min-h-screen flex items-center justify-center bg-gray-900">
						<div className="w-full max-w-md p-8 rounded-lg bg-gray-800">
							<form onSubmit={handleSubmit} className="space-y-4">
								<div>
									<label htmlFor="projectName" className="block text-sm font-medium text-gray-200 mb-1">
										Project Name
									</label>
									<Input
										type="text"
										id="projectName"
										value={projectName}
										onChange={(e) => setProjectName(e.target.value)}
										placeholder="Enter project name"
										className="w-full bg-gray-700 text-white placeholder-gray-400 border-gray-600 focus:border-blue-500 focus:ring-blue-500"
										required
									/>
								</div>
								<Button type="submit" className="w-full bg-blue-600 hover:bg-blue-700 text-white">
									Submit
								</Button>
							</form>
						</div>
					</div>
				</>
			}
		</>
	)

}

export default CreateNewProject
