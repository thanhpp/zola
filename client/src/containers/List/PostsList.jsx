import React from "react";
import Posts from "../../components/list/Posts";

const listData = [];
for (let i = 0; i < 23; i++) {
	listData.push({
		id: i,
		author: "John Doe",
		avatar: "https://joeschmoe.io/api/v1/random",
		media: [
			"https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png",
			"https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png",
			"https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png",
		],
		content:
			"We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure),\
            to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure),\
            to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure),\
            to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure),\
            to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure),\
            to help people create their product prototypes beautifully and efficiently.",
		like: 126,
		comment: 126,
	});
}

export default function PostsList() {
	return <Posts posts={listData} />;
}
