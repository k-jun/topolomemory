package models

type Card struct {
	Id            int
	Name          string
	Svg           string
	TopologyHoles int
	TopologyParts int
}

var deck []Card
var field []Card

func init() {
	deck = []Card{
		// 1/0
		{
			Id:            1,
			Name:          "clip",
			TopologyHoles: 1,
			TopologyParts: 0,
			Svg:           `<svg width="80px" height="80px" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-labelledby="clipIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="clipIconTitle">Attachment (paper clip)</title> <path d="M7.93517 13.7796L15.1617 6.55304C16.0392 5.67631 17.4657 5.67631 18.3432 6.55304C19.2206 7.43052 19.2206 8.85774 18.3432 9.73522L8.40091 19.5477C6.9362 21.0124 4.56325 21.0124 3.09854 19.5477C1.63382 18.0837 1.63382 15.7093 3.09854 14.2453L12.9335 4.53784C14.984 2.48739 18.3094 2.48739 20.3569 4.53784C22.4088 6.58904 22.4088 9.91146 20.3584 11.9619L13.239 19.082"/></svg>`,
		},
		{
			Id:            2,
			Name:          "dolar",
			TopologyHoles: 1,
			TopologyParts: 0,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="dolarIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="dolarIconTitle">Dolar</title> <path d="M12 4L12 6M12 18L12 20M15.5 8C15.1666667 6.66666667 14 6 12 6 9 6 8.5 7.95652174 8.5 9 8.5 13.140327 15.5 10.9649412 15.5 15 15.5 16.0434783 15 18 12 18 10 18 8.83333333 17.3333333 8.5 16"/> </svg>`,
		},
		{
			Id:            3,
			Name:          "eye-closed",
			TopologyHoles: 1,
			TopologyParts: 0,
			Svg:           `<svg width="80px" height="80px" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-labelledby="eyeClosedIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="eyeClosedIconTitle">Hidden (closed eye)</title> <path d="M20 9C20 9 19.6797 9.66735 19 10.5144M12 14C10.392 14 9.04786 13.5878 7.94861 13M12 14C13.608 14 14.9521 13.5878 16.0514 13M12 14V17.5M4 9C4 9 4.35367 9.73682 5.10628 10.6448M7.94861 13L5 16M7.94861 13C6.6892 12.3266 5.75124 11.4228 5.10628 10.6448M16.0514 13L18.5 16M16.0514 13C17.3818 12.2887 18.3535 11.3202 19 10.5144M5.10628 10.6448L2 12M19 10.5144L22 12"/> </svg>`,
		},
		{
			Id:            4,
			Name:          "hearing-disability",
			TopologyHoles: 1,
			TopologyParts: 0,
			Svg:           `<svg width="80px" height="80px" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-labelledby="hearingDisabilityIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="hearingDisabilityIconTitle">Hearing disability</title> <path d="M6 11.2632C6 7.80411 8.80411 5 12.2632 5C15.7222 5 18.5263 7.80411 18.5263 11.2632C18.5263 13.833 16.9197 15.8089 15.3947 17.5263C14.1124 18.9705 13.6053 22 10.4737 22C9.46637 22 8.53681 21.6671 7.78904 21.1053"/> <path d="M9.57892 11.2631C9.57892 9.78068 10.7807 8.57892 12.2631 8.57892C13.7456 8.57892 14.9473 9.78068 14.9473 11.2631C14.9473 11.9506 14.6889 12.5777 14.2639 13.0526"/> <path d="M3 21L21 3"/> </svg>`,
		},
		{
			Id:            5,
			Name:          "people",
			TopologyHoles: 1,
			TopologyParts: 0,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="peopleIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="peopleIconTitle">People</title> <path d="M1 18C1 15.75 4 15.75 5.5 14.25 6.25 13.5 4 13.5 4 9.75 4 7.25025 4.99975 6 7 6 9.00025 6 10 7.25025 10 9.75 10 13.5 7.75 13.5 8.5 14.25 10 15.75 13 15.75 13 18M12.7918114 15.7266684C13.2840551 15.548266 13.6874862 15.3832994 14.0021045 15.2317685 14.552776 14.9665463 15.0840574 14.6659426 15.5 14.25 16.25 13.5 14 13.5 14 9.75 14 7.25025 14.99975 6 17 6 19.00025 6 20 7.25025 20 9.75 20 13.5 17.75 13.5 18.5 14.25 20 15.75 23 15.75 23 18"/> <path stroke-linecap="round" d="M12,16 C12.3662741,15.8763472 12.6302112,15.7852366 12.7918114,15.7266684"/> </svg>`,
		},
		{
			Id:            6,
			Name:          "shuffle",
			TopologyHoles: 1,
			TopologyParts: 0,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="shuffleIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="shuffleIconTitle">Shuffle</title> <path d="M21 16.0399H17.7707C15.8164 16.0399 13.9845 14.9697 12.8611 13.1716L10.7973 9.86831C9.67384 8.07022 7.84196 7 5.88762 7L3 7"/> <path d="M21 7H17.7707C15.8164 7 13.9845 8.18388 12.8611 10.1729L10.7973 13.8271C9.67384 15.8161 7.84196 17 5.88762 17L3 17"/> <path d="M19 4L22 7L19 10"/> <path d="M19 13L22 16L19 19"/> </svg>`,
		},
		// 1/1
		{
			Id:            7,
			Name:          "anchor",
			TopologyHoles: 1,
			TopologyParts: 1,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="anchorIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="anchorIconTitle">Anchor</title> <path d="M12 20L12 7M9 10L15 10"/> <circle cx="12" cy="5" r="2"/> <path d="M20,14 C18.6666667,18 16,20 12,20 C8,20 5.33333333,18 4,14"/> </svg>`,
		},
		{
			Id:            8,
			Name:          "hourglass",
			TopologyHoles: 1,
			TopologyParts: 1,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="hourglassIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="hourglassIconTitle">Hourglass</title> <path d="M16.7647059,7.82108056 L16.7647059,5 C16.7647059,3.8954305 15.8692754,3 14.7647059,3 L9.23529412,3 C8.13072462,3 7.23529412,3.8954305 7.23529412,5 L7.23529412,7.82108056 L11.4142136,12 L7.23529412,16.1789194 L7.23529412,19 C7.23529412,20.1045695 8.13072462,21 9.23529412,21 L14.7647059,21 C15.8692754,21 16.7647059,20.1045695 16.7647059,19 L16.7647059,16.1789194 L12.5857864,12 L16.7647059,7.82108056 Z"/> </svg>`,
		},
		{
			Id:            9,
			Name:          "search",
			TopologyHoles: 1,
			TopologyParts: 1,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="searchIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="searchIconTitle">Search</title> <path d="M14.4121122,14.4121122 L20,20"/> <circle cx="10" cy="10" r="6"/> </svg>`,
		},
		{
			Id:            10,
			Name:          "share-ios",
			TopologyHoles: 1,
			TopologyParts: 1,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="shareiOSIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="shareiOSIconTitle">Share</title> <path d="M12,3 L12,15"/> <polyline points="9 5 12 2 15 5"/> <rect width="12" height="12" x="6" y="9"/> </svg>`,
		},
		{
			Id:            11,
			Name:          "tool",
			TopologyHoles: 1,
			TopologyParts: 1,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="toolIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="toolIconTitle">Tool</title> <path d="M9.74292939,13.7429294 C9.19135019,13.9101088 8.60617271,14 8,14 C4.6862915,14 2,11.3137085 2,8 C2,7.07370693 2.20990431,6.19643964 2.58474197,5.4131691 L6.94974747,9.77817459 L9.77817459,6.94974747 L5.4131691,2.58474197 C6.19643964,2.20990431 7.07370693,2 8,2 C11.3137085,2 14,4.6862915 14,8 C14,8.88040772 13.8103765,9.71652648 13.4697429,10.4697429 L20.5858636,17.5858636 C21.3669122,18.3669122 21.3669122,19.6332422 20.5858636,20.4142907 L19.9142907,21.0858636 C19.1332422,21.8669122 17.8669122,21.8669122 17.0858636,21.0858636 L9.74292939,13.7429294 Z"/> </svg>`,
		},
		{
			Id:            12,
			Name:          "wine",
			TopologyHoles: 1,
			TopologyParts: 1,
			Svg:           `<svg width="80px" height="80px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" aria-labelledby="wineIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" color="#2329D6"> <title id="wineIconTitle">Wine</title> <path d="M12 14C14.7614 14 17 11.7614 17 9C17 7.07098 15.4615 4 15.4615 4H8.53846C8.53846 4 7 6.23858 7 9C7 11.7614 9.23858 14 12 14ZM12 14V20M12 20H8M12 20H16"/> </svg>`,
		},
		// 1/3
		{
			Id:            13,
			Name:          "contact-book",
			TopologyHoles: 1,
			TopologyParts: 3,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="contactBookIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="contactBookIconTitle">Contact Book</title> <path d="M8,3 L8,21"/> <path d="M20,2.99999899 L20,21 L6,21 C4.8954305,21 4,20.1045695 4,19 L4,4.99999899 C4,3.89542949 4.8954305,2.99999899 6,2.99999899 L20,2.99999899 Z"/> <path d="M8,17 C8,15 11.3333333,15.3333333 12.6666667,14 C13.3333333,13.3333333 11.3333333,13.3333333 11.3333333,10 C11.3333333,7.778 12.222,6.66666667 14,6.66666667 C15.778,6.66666667 16.6666667,7.778 16.6666667,10 C16.6666667,13.3333333 14.6666667,13.3333333 15.3333333,14 C16.6666667,15.3333333 20,15 20,17"/> </svg>`,
		},
		{
			Id:            14,
			Name:          "horn",
			TopologyHoles: 1,
			TopologyParts: 3,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="hornIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="hornIconTitle">Bullhorn</title> <path stroke-linejoin="round" d="M6.5,9 C8.33333333,9 10.1666667,9 12,9 C14,9 16.3333333,7.33333333 19,4 L19,19 C16.3333333,15.6666667 14,14 12,14 C10.1666667,14 8.33333333,14 6.5,14 L6.5,14 C5.11928813,14 4,12.8807119 4,11.5 L4,11.5 C4,10.1192881 5.11928813,9 6.5,9 Z"/> <polygon points="7 14 9 20 13 20 11 14"/> <path d="M11,9 L11,14"/> </svg>`,
		},
		{
			Id:            15,
			Name:          "layout-left",
			TopologyHoles: 1,
			TopologyParts: 3,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="layoutLeftIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="layoutLeftIconTitle">Layout Left</title> <rect width="18" height="18" x="3" y="3"/> <path d="M3 8L21 8M9 8L9 21"/> </svg>`,
		},
		{
			Id:            16,
			Name:          "map",
			TopologyHoles: 1,
			TopologyParts: 3,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="mapIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="mapIconTitle">Map</title> <polygon points="9 19 3 21 3 5 9 3 15 5 21 3 21 18.5 15 21"/> <path stroke-linecap="round" d="M15 5L15 21M9 3L9 19"/> </svg>`,
		},
		{
			Id:            17,
			Name:          "poll",
			TopologyHoles: 1,
			TopologyParts: 3,
			Svg:           `<svg width="80px" height="80px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" aria-labelledby="pollIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" color="#2329D6"> <title id="pollIconTitle">Poll</title> <path d="M4 4V20"/> <path d="M4 6H15V10H4"/> <path d="M4 10H19V14H4"/> <path d="M4 14H12V18H4"/> </svg>`,
		},
		{
			Id:            18,
			Name:          "save",
			TopologyHoles: 1,
			TopologyParts: 3,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="saveIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="saveIconTitle">Save</title> <path d="M17.2928932,3.29289322 L21,7 L21,20 C21,20.5522847 20.5522847,21 20,21 L4,21 C3.44771525,21 3,20.5522847 3,20 L3,4 C3,3.44771525 3.44771525,3 4,3 L16.5857864,3 C16.8510029,3 17.1053568,3.10535684 17.2928932,3.29289322 Z"/> <rect width="10" height="8" x="7" y="13"/> <rect width="8" height="5" x="8" y="3"/> </svg>`,
		},
		// 2/0
		{
			Id:            19,
			Name:          "bin",
			TopologyHoles: 2,
			TopologyParts: 0,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="binIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="binIconTitle">Bin</title> <path d="M19 6L5 6M14 5L10 5M6 10L6 20C6 20.6666667 6.33333333 21 7 21 7.66666667 21 11 21 17 21 17.6666667 21 18 20.6666667 18 20 18 19.3333333 18 16 18 10"/> </svg>`,
		},
		{
			Id:            20,
			Name:          "chevrons-down",
			TopologyHoles: 2,
			TopologyParts: 0,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="chevronsDownIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="chevronsDownIconTitle">Chevrons Down</title> <polyline points="7 13 12 18 17 13"/> <polyline points="7 7 12 12 17 7"/> </svg>`,
		},
		{
			Id:            21,
			Name:          "ear",
			TopologyHoles: 2,
			TopologyParts: 0,
			Svg:           `<svg width="80px" height="80px" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-labelledby="earIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="earIconTitle">Ear (hearing)</title> <path d="M6 10C6 6.13401 9.13401 3 13 3C16.866 3 20 6.13401 20 10C20 12.8721 18.2043 15.0806 16.5 17C15.0668 18.6141 14.5 22 11 22C9.87418 22 8.83526 21.6279 7.99951 21"/> <path d="M10 10C10 8.34315 11.3431 7 13 7C14.6569 7 16 8.34315 16 10C16 10.7684 15.7111 11.4692 15.2361 12"/> </svg>`,
		},
		{
			Id:            22,
			Name:          "history",
			TopologyHoles: 2,
			TopologyParts: 0,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="historyIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="historyIconTitle">History</title> <polyline points="1 12 3 14 5 12"/> <polyline points="12 7 12 12 15 15"/> <path d="M12,21 C16.9705627,21 21,16.9705627 21,12 C21,7.02943725 16.9705627,3 12,3 C7.02943725,3 3,7.02943725 3,12 C3,11.975305 3,12.3086383 3,13"/> </svg>`,
		},
		{
			Id:            23,
			Name:          "rotate",
			TopologyHoles: 2,
			TopologyParts: 0,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="rotateIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="rotateIconTitle">Rotate</title> <path d="M22 12l-3 3-3-3"/> <path d="M2 12l3-3 3 3"/> <path d="M19.016 14v-1.95A7.05 7.05 0 0 0 8 6.22"/> <path d="M16.016 17.845A7.05 7.05 0 0 1 5 12.015V10"/> <path stroke-linecap="round" d="M5 10V9"/> <path stroke-linecap="round" d="M19 15v-1"/> </svg>`,
		},
		{
			Id:            24,
			Name:          "thunder",
			TopologyHoles: 2,
			TopologyParts: 0,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="thunderIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="thunderIconTitle">Thunder</title> <polyline points="13 15 14 18 12 18 13 21"/> <path d="M19.051177,17.9568482 C20.5,17.9709234 22,16.2454604 22,14.5 C22,12.7095527 20.6555928,11.2331085 18.9211951,11.0250841 C18.4554927,8.17503894 15.9817502,6 13,6 C10.711801,6 8.72277,7.28089089 7.71081142,9.16476838 C7.3255638,9.05739789 6.9194849,9 6.5,9 C4.01471863,9 2,11.0147186 2,13.5 C2,15.8113421 3.5,17.9709234 5.98562648,17.9709234 C5.98562648,17.9709234 6.32375099,17.9709234 7,17.9709234"/> </svg>`,
		},
		// 2/1
		{
			Id:            25,
			Name:          "thunder",
			TopologyHoles: 2,
			TopologyParts: 1,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="thunderIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="thunderIconTitle">Thunder</title> <polyline points="13 15 14 18 12 18 13 21"/> <path d="M19.051177,17.9568482 C20.5,17.9709234 22,16.2454604 22,14.5 C22,12.7095527 20.6555928,11.2331085 18.9211951,11.0250841 C18.4554927,8.17503894 15.9817502,6 13,6 C10.711801,6 8.72277,7.28089089 7.71081142,9.16476838 C7.3255638,9.05739789 6.9194849,9 6.5,9 C4.01471863,9 2,11.0147186 2,13.5 C2,15.8113421 3.5,17.9709234 5.98562648,17.9709234 C5.98562648,17.9709234 6.32375099,17.9709234 7,17.9709234"/> </svg>`,
		},
		{
			Id:            26,
			Name:          "cancel",
			TopologyHoles: 2,
			TopologyParts: 1,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="cancelIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="cancelIconTitle">Cancel</title> <path d="M15.5355339 15.5355339L8.46446609 8.46446609M15.5355339 8.46446609L8.46446609 15.5355339"/> <path d="M4.92893219,19.0710678 C1.02368927,15.1658249 1.02368927,8.83417511 4.92893219,4.92893219 C8.83417511,1.02368927 15.1658249,1.02368927 19.0710678,4.92893219 C22.9763107,8.83417511 22.9763107,15.1658249 19.0710678,19.0710678 C15.1658249,22.9763107 8.83417511,22.9763107 4.92893219,19.0710678 Z"/> </svg>`,
		},
		{
			Id:            27,
			Name:          "restaurant",
			TopologyHoles: 2,
			TopologyParts: 1,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="restaurantIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="restaurantIconTitle">Restaurant</title> <path d="M8 4L8 20M18 12C16 11.3333333 15 10 15 8 15 6 16 4.66666667 18 4L18 20 18 12zM5 4L5 7C5 8.65685425 6.34314575 10 8 10L8 10C9.65685425 10 11 8.65685425 11 7L11 4"/> </svg>`,
		},
		{
			Id:            28,
			Name:          "shift",
			TopologyHoles: 2,
			TopologyParts: 1,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="shiftIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="shiftIconTitle">Shift</title> <path d="M5,21 L19,21 L5,21 Z M16,12 L16,17 L8,17 L8,12 L3,12 L12,3 L21,12 L16,12 Z"/> </svg>`,
		},
		{
			Id:            29,
			Name:          "thumb-up",
			TopologyHoles: 2,
			TopologyParts: 1,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="thumbUpIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="thumbUpIconTitle">Thumb Up</title> <path d="M8,8.73984815 C8,8.26242561 8.17078432,7.80075162 8.4814868,7.43826541 L13.2723931,1.84887469 C13.7000127,1.34998522 14.4122932,1.20614658 15,1.5 C15.5737957,1.78689785 15.849314,2.45205792 15.6464466,3.06066017 L14,8 L18.6035746,8 C18.7235578,8 18.8432976,8.01079693 18.9613454,8.03226018 C20.0480981,8.22985158 20.7689058,9.27101818 20.5713144,10.3577709 L19.2985871,17.3577709 C19.1256814,18.3087523 18.2974196,19 17.3308473,19 L10,19 C8.8954305,19 8,18.1045695 8,17 L8,8.73984815 Z"/> <path d="M4,18 L4,9"/> </svg>`,
		},
		{
			Id:            30,
			Name:          "time",
			TopologyHoles: 2,
			TopologyParts: 1,
			Svg:           `<svg role="img" xmlns="http://www.w3.org/2000/svg" width="80px" height="80px" viewBox="0 0 24 24" aria-labelledby="timeIconTitle" stroke="#2329D6" stroke-width="0.6" stroke-linecap="square" stroke-linejoin="miter" fill="none" color="#2329D6"> <title id="timeIconTitle">Time</title> <circle cx="12" cy="12" r="10"/> <polyline points="12 5 12 12 16 16"/> </svg>`,
		},
	}
}
