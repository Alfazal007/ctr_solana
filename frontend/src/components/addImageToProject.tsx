import { UserContext } from '@/context/UserContext';
import { useContext, useEffect, useRef, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import axios from "axios"
import { toast } from '@/hooks/use-toast';
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { Upload } from 'lucide-react'
import Navbar from './Navbar';

export interface Project {
	name: string
	id: string
	creatorId: string
	started: boolean
	completed: boolean
}

const AddImageToProject = () => {
	const { projectId } = useParams();
	const { user } = useContext(UserContext)
	const navigate = useNavigate()
	const [project, setProject] = useState<Project | null>(null)
	const [selectedFile, setSelectedFile] = useState<File | null>(null)
	const fileInputRef = useRef<HTMLInputElement>(null)

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
	}

	useEffect(() => {
		if (!user) {
			navigate("/signin")
			return
		}
		fetchProject()
	}, [])



	const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		if (e.target.files && e.target.files[0]) {
			setSelectedFile(e.target.files[0])
		}
	}

	async function uploadFile() {
		try {

			if (!selectedFile) {
				toast({ title: "Select something to upload" })
				return
			}
			const formData = new FormData();
			formData.append('image', selectedFile);
			const uploadImageResponse = await axios.post(`http://localhost:8000/api/v1/project/add-image/${projectId}`, {
				image: selectedFile
			}, {
				withCredentials: true, headers: {
					'Content-Type': 'multipart/form-data'
				}
			})
			if (uploadImageResponse.status != 200) {
				toast({ title: "Issue uploading the image", variant: "destructive" })
			} else {
				toast({ title: "Uploaded image successfully" })
			}
		} catch (err) {
			toast({ title: "Issue uploading the image", variant: "destructive" })
		} finally {
			setSelectedFile(null)
			if (fileInputRef.current) {
				fileInputRef.current.value = ""
			}
		}
	}

	const handleSubmit = (e: React.FormEvent) => {
		e.preventDefault()
		if (selectedFile) {
			uploadFile()
		}
	}

	async function startProject() {
		if (!project) {
			toast({
				title: "Project not found",
				variant: "destructive"
			})
			return
		}
		try {
			const startProjectResponse = await axios.put(`http://localhost:8000/api/v1/project/start/${project.id}`, {}, { withCredentials: true })
			if (startProjectResponse.status != 200) {
				toast({ title: "issue starting the project", variant: "destructive" })
			} else {
				toast({ title: "Started the project voting" })
				navigate("/")
			}
		} catch (err) {
			toast({ title: "issue starting the project", variant: "destructive" })
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
							<form onSubmit={handleSubmit} className="space-y-6">
								<div>
									<Label htmlFor="imageUpload" className="block text-lg font-medium text-gray-200 mb-2">
										Upload Image
									</Label>
									<div className="mt-1 flex justify-center px-6 pt-5 pb-6 border-2 border-gray-600 border-dashed rounded-md">
										<div className="space-y-1 text-center">
											<Upload className="mx-auto h-12 w-12 text-gray-400" />
											<div className="flex text-sm text-gray-400">
												<label
													htmlFor="imageUpload"
													className="relative cursor-pointer rounded-md font-medium text-blue-500 hover:text-blue-400 focus-within:outline-none focus-within:ring-2 focus-within:ring-offset-2 focus-within:ring-blue-500"
												>
													<span className="ml-5">Upload a file</span>
													<Input
														id="imageUpload"
														name="imageUpload"
														type="file"
														accept="image/*"
														className="sr-only"
														onChange={handleFileChange}
														ref={fileInputRef}
													/>
												</label>
											</div>
											<p className="text-xs text-gray-400">PNG or JPG up to 50kb</p>
										</div>
									</div>
								</div>
								{selectedFile && (
									<p className="text-sm text-gray-300">
										Selected file: {selectedFile.name}
									</p>
								)}
								<Button
									type="submit"
									className="w-full bg-blue-600 hover:bg-blue-700 text-white"
									disabled={!selectedFile}
								>
									Upload Image
								</Button>
								<Button
									onClick={startProject}
									disabled={project?.started}
									className="w-full bg-blue-600 hover:bg-blue-700 text-white"
								>
									Start project
								</Button>
							</form>
						</div>
					</div>
				</>
			}
		</>
	)
}

export default AddImageToProject
