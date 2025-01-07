interface NumberedImageProps {
	src: string;
	alt: string;
	number: number;
}

function NumberedImage({ src, alt, number }: NumberedImageProps) {
	return (
		<div className="relative">
			<img
				src={src}
				alt={alt}
				className="w-full h-auto rounded-lg object-cover"
			/>
			<div className="absolute top-2 right-2 bg-blue-900 text-white rounded-full w-8 h-8 flex items-center justify-center font-bold">
				{number}
			</div>
		</div>
	);
}

export default NumberedImage;

