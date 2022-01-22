import { useState, useEffect } from "react";

const useResize = (myRef) => {
	const [width, setWidth] = useState(0);
	const [height, setHeight] = useState(0);

	useEffect(() => {
		console.log("re rendering");
		const handleResize = () => {
			setWidth(myRef.current.offsetWidth);
			setHeight(myRef.current.offsetHeight);
		};

		const ref = myRef.current;
		ref && myRef.current.addEventListener("resize", handleResize);

		return () => {
			ref.removeEventListener("resize", handleResize);
		};
	}, [myRef]);

	return { width, height };
};

export default useResize;
