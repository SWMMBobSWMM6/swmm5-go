// Generated from EPA SWMM 5.2.4 (xsect.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var N_A_Circ int32 = 51
var A_Circ = [51]float64{float64(0.0), float64(0.00471), float64(0.0134), float64(0.024446), float64(0.0374), float64(0.05208), float64(0.0680), float64(0.08505), float64(0.1033), float64(0.12236), float64(0.1423), float64(0.16310), float64(0.1845), float64(0.20665), float64(0.2292), float64(0.25236), float64(0.2759), float64(0.29985), float64(0.3242), float64(0.34874), float64(0.3736), float64(0.39878), float64(0.4237), float64(0.44907), float64(0.4745), float64(0.500), float64(0.5255), float64(0.55093), float64(0.5763), float64(0.60135), float64(0.6264), float64(0.65126), float64(0.6758), float64(0.70015), float64(0.7241), float64(0.74764), float64(0.7708), float64(0.79335), float64(0.8154), float64(0.83690), float64(0.8576), float64(0.87764), float64(0.8967), float64(0.91495), float64(0.9320), float64(0.94792), float64(0.9626), float64(0.97555), float64(0.9866), float64(0.99516), float64(1.000)}
var N_R_Circ int32 = 51
var R_Circ = [51]float64{float64(0.0100), float64(0.0528), float64(0.1048), float64(0.1556), float64(0.2052), float64(0.254), float64(0.3016), float64(0.3484), float64(0.3944), float64(0.4388), float64(0.4824), float64(0.5248), float64(0.5664), float64(0.6064), float64(0.6456), float64(0.6836), float64(0.7204), float64(0.7564), float64(0.7912), float64(0.8244), float64(0.8568), float64(0.888), float64(0.9176), float64(0.9464), float64(0.9736), float64(1.00), float64(1.024), float64(1.048), float64(1.070), float64(1.0912), float64(1.110), float64(1.1272), float64(1.144), float64(1.1596), float64(1.174), float64(1.1848), float64(1.194), float64(1.2024), float64(1.210), float64(1.2148), float64(1.217), float64(1.2172), float64(1.215), float64(1.2104), float64(1.203), float64(1.192), float64(1.178), float64(1.1584), float64(1.132), float64(1.094), float64(1.000)}
var N_Y_Circ int32 = 51
var Y_Circ = [51]float64{float64(0.0), float64(0.05236), float64(0.08369), float64(0.11025), float64(0.13423), float64(0.15643), float64(0.17755), float64(0.19772), float64(0.21704), float64(0.23581), float64(0.25412), float64(0.27194), float64(0.28948), float64(0.30653), float64(0.32349), float64(0.34017), float64(0.35666), float64(0.37298), float64(0.38915), float64(0.40521), float64(0.42117), float64(0.43704), float64(0.45284), float64(0.46858), float64(0.4843), float64(0.50000), float64(0.51572), float64(0.53146), float64(0.54723), float64(0.56305), float64(0.57892), float64(0.59487), float64(0.61093), float64(0.62710), float64(0.64342), float64(0.65991), float64(0.67659), float64(0.69350), float64(0.71068), float64(0.72816), float64(0.74602), float64(0.76424), float64(0.78297), float64(0.80235), float64(0.82240), float64(0.84353), float64(0.86563), float64(0.88970), float64(0.91444), float64(0.94749), float64(1.0)}
var N_S_Circ int32 = 51
var S_Circ = [51]float64{float64(0.0), float64(0.00529), float64(0.01432), float64(0.02559), float64(0.03859), float64(0.05304), float64(0.06877), float64(0.08551), float64(0.10326), float64(0.12195), float64(0.14144), float64(0.16162), float64(0.18251), float64(0.2041), float64(0.22636), float64(0.24918), float64(0.27246), float64(0.29614), float64(0.32027), float64(0.34485), float64(0.36989), float64(0.39531), float64(0.42105), float64(0.44704), float64(0.47329), float64(0.4998), float64(0.52658), float64(0.55354), float64(0.58064), float64(0.60777), float64(0.63499), float64(0.66232), float64(0.68995), float64(0.7177), float64(0.74538), float64(0.77275), float64(0.79979), float64(0.82658), float64(0.8532), float64(0.87954), float64(0.90546), float64(0.93095), float64(0.95577), float64(0.97976), float64(1.00291), float64(1.02443), float64(1.04465), float64(1.06135), float64(1.08208), float64(1.07662), float64(1.0)}
var N_W_Circ int32 = 51
var W_Circ = [51]float64{float64(0.0), float64(0.2800), float64(0.3919), float64(0.4750), float64(0.5426), float64(0.6000), float64(0.6499), float64(0.6940), float64(0.7332), float64(0.7684), float64(0.8000), float64(0.8285), float64(0.8542), float64(0.8773), float64(0.8980), float64(0.9165), float64(0.9330), float64(0.9474), float64(0.9600), float64(0.9708), float64(0.9798), float64(0.9871), float64(0.9928), float64(0.9968), float64(0.9992), float64(1.000), float64(0.9992), float64(0.9968), float64(0.9928), float64(0.9871), float64(0.9798), float64(0.9708), float64(0.9600), float64(0.9474), float64(0.9330), float64(0.9165), float64(0.8980), float64(0.8773), float64(0.8542), float64(0.8285), float64(0.8000), float64(0.7684), float64(0.7332), float64(0.6940), float64(0.6499), float64(0.6000), float64(0.5426), float64(0.4750), float64(0.3919), float64(0.2800), float64(0.0)}
var N_A_Egg int32 = 26
var A_Egg = [26]float64{float64(0.0000), float64(0.0116), float64(0.0316), float64(0.0550), float64(0.0850), float64(0.1200), float64(0.1555), float64(0.1900), float64(0.2250), float64(0.2750), float64(0.3200), float64(0.3700), float64(0.4200), float64(0.4700), float64(0.5150), float64(0.5700), float64(0.6200), float64(0.6800), float64(0.7300), float64(0.7800), float64(0.8350), float64(0.8850), float64(0.9250), float64(0.9550), float64(0.9800), float64(1.000)}
var N_R_Egg int32 = 26
var R_Egg = [26]float64{float64(0.0100), float64(0.1300), float64(0.2440), float64(0.3020), float64(0.3860), float64(0.4650), float64(0.5360), float64(0.6110), float64(0.6760), float64(0.7350), float64(0.7910), float64(0.8540), float64(0.9040), float64(0.9410), float64(1.008), float64(1.045), float64(1.076), float64(1.115), float64(1.146), float64(1.162), float64(1.186), float64(1.193), float64(1.186), float64(1.162), float64(1.107), float64(1.000)}
var N_Y_Egg int32 = 51
var Y_Egg = [51]float64{float64(0.0), float64(0.04912), float64(0.08101), float64(0.11128), float64(0.14161), float64(0.16622), float64(0.18811), float64(0.21356), float64(0.23742), float64(0.25742), float64(0.27742), float64(0.29741), float64(0.31742), float64(0.33742), float64(0.35747), float64(0.37364), float64(0.4), float64(0.41697), float64(0.43372), float64(0.45), float64(0.46374), float64(0.47747), float64(0.49209), float64(0.50989), float64(0.53015), float64(0.55), float64(0.56429), float64(0.57675), float64(0.58834), float64(0.6), float64(0.61441), float64(0.62967), float64(0.64582), float64(0.66368), float64(0.68209), float64(0.7), float64(0.71463), float64(0.72807), float64(0.74074), float64(0.75296), float64(0.765), float64(0.77784), float64(0.79212), float64(0.80945), float64(0.82936), float64(0.85), float64(0.86731), float64(0.88769), float64(0.914), float64(0.95), float64(1.0)}
var N_S_Egg int32 = 51
var S_Egg = [51]float64{float64(0.0), float64(0.00295), float64(0.01331), float64(0.02629), float64(0.04), float64(0.05657), float64(0.075), float64(0.09432), float64(0.11473), float64(0.13657), float64(0.15894), float64(0.1803), float64(0.20036), float64(0.22), float64(0.23919), float64(0.25896), float64(0.28), float64(0.30504), float64(0.33082), float64(0.35551), float64(0.37692), float64(0.39809), float64(0.42), float64(0.44625), float64(0.47321), float64(0.5), float64(0.52255), float64(0.54481), float64(0.56785), float64(0.59466), float64(0.62485), float64(0.65518), float64(0.68181), float64(0.70415), float64(0.72585), float64(0.74819), float64(0.77482), float64(0.80515), float64(0.83534), float64(0.86193), float64(0.88465), float64(0.9069), float64(0.93), float64(0.95866), float64(0.98673), float64(1.01238), float64(1.03396), float64(1.05), float64(1.06517), float64(1.0538), float64(1.0)}
var N_W_Egg int32 = 26
var W_Egg = [26]float64{float64(0.0), float64(0.3250), float64(0.4270), float64(0.5080), float64(0.5820), float64(0.6420), float64(0.6960), float64(0.7460), float64(0.7910), float64(0.8360), float64(0.8660), float64(0.8960), float64(0.9260), float64(0.9560), float64(0.9700), float64(0.9850), float64(1.000), float64(0.9850), float64(0.9700), float64(0.9400), float64(0.8960), float64(0.8360), float64(0.7640), float64(0.6420), float64(0.3100), float64(0.0)}
var N_A_Horseshoe int32 = 26
var A_Horseshoe = [26]float64{float64(0.0000), float64(0.0181), float64(0.0508), float64(0.0908), float64(0.1326), float64(0.1757), float64(0.2201), float64(0.2655), float64(0.3118), float64(0.3587), float64(0.4064), float64(0.4542), float64(0.5023), float64(0.5506), float64(0.5987), float64(0.6462), float64(0.6931), float64(0.7387), float64(0.7829), float64(0.8253), float64(0.8652), float64(0.9022), float64(0.9356), float64(0.9645), float64(0.9873), float64(1.000)}
var N_R_Horseshoe int32 = 26
var R_Horseshoe = [26]float64{float64(0.0100), float64(0.1040), float64(0.2065), float64(0.3243), float64(0.4322), float64(0.5284), float64(0.6147), float64(0.6927), float64(0.7636), float64(0.8268), float64(0.8873), float64(0.9417), float64(0.9905), float64(1.036), float64(1.077), float64(1.113), float64(1.143), float64(1.169), float64(1.189), float64(1.202), float64(1.208), float64(1.206), float64(1.195), float64(1.170), float64(1.126), float64(1.000)}
var N_Y_Horseshoe int32 = 51
var Y_Horseshoe = [51]float64{float64(0.0), float64(0.04146), float64(0.07033), float64(0.09098), float64(0.10962), float64(0.12921), float64(0.14813), float64(0.16701), float64(0.18565), float64(0.20401), float64(0.22211), float64(0.23998), float64(0.25769), float64(0.27524), float64(0.29265), float64(0.3099), float64(0.32704), float64(0.34406), float64(0.36101), float64(0.3779), float64(0.39471), float64(0.41147), float64(0.42818), float64(0.44484), float64(0.46147), float64(0.47807), float64(0.49468), float64(0.51134), float64(0.52803), float64(0.54474), float64(0.56138), float64(0.57804), float64(0.59478), float64(0.61171), float64(0.62881), float64(0.64609), float64(0.6635), float64(0.68111), float64(0.69901), float64(0.71722), float64(0.73583), float64(0.7549), float64(0.77447), float64(0.79471), float64(0.81564), float64(0.83759), float64(0.86067), float64(0.88557), float64(0.91159), float64(0.9452), float64(1.0)}
var N_S_Horseshoe int32 = 51
var S_Horseshoe = [51]float64{float64(0.0), float64(0.00467), float64(0.01237), float64(0.02268), float64(0.03515), float64(0.04943), float64(0.06525), float64(0.08212), float64(0.10005), float64(0.11891), float64(0.13856), float64(0.15896), float64(0.18004), float64(0.20172), float64(0.22397), float64(0.24677), float64(0.27006), float64(0.2938), float64(0.3179), float64(0.34237), float64(0.3672), float64(0.39239), float64(0.41792), float64(0.44374), float64(0.46984), float64(0.49619), float64(0.52276), float64(0.5495), float64(0.5764), float64(0.60345), float64(0.63065), float64(0.65795), float64(0.68531), float64(0.71271), float64(0.74009), float64(0.76738), float64(0.79451), float64(0.82144), float64(0.84814), float64(0.8745), float64(0.90057), float64(0.92652), float64(0.95244), float64(0.97724), float64(0.99988), float64(1.02048), float64(1.03989), float64(1.05698), float64(1.07694), float64(1.07562), float64(1.0)}
var N_W_Horseshoe int32 = 26
var W_Horseshoe = [26]float64{float64(0.0), float64(0.5878), float64(0.8772), float64(0.8900), float64(0.9028), float64(0.9156), float64(0.9284), float64(0.9412), float64(0.9540), float64(0.9668), float64(0.9798), float64(0.9928), float64(0.9992), float64(0.9992), float64(0.9928), float64(0.9798), float64(0.9600), float64(0.9330), float64(0.8980), float64(0.8542), float64(0.8000), float64(0.7332), float64(0.6499), float64(0.5426), float64(0.3919), float64(0.0)}
var N_Y_Gothic int32 = 51
var Y_Gothic = [51]float64{float64(0.0), float64(0.04522), float64(0.07825), float64(0.10646), float64(0.12645), float64(0.14645), float64(0.16787), float64(0.18641), float64(0.20129), float64(0.22425), float64(0.24129), float64(0.25624), float64(0.27344), float64(0.29097), float64(0.30529), float64(0.32607), float64(0.33755), float64(0.35073), float64(0.36447), float64(0.37558), float64(0.4), float64(0.4181), float64(0.43648), float64(0.45374), float64(0.46805), float64(0.48195), float64(0.49626), float64(0.51352), float64(0.5319), float64(0.55), float64(0.56416), float64(0.57787), float64(0.59224), float64(0.6095), float64(0.62941), float64(0.65), float64(0.67064), float64(0.69055), float64(0.70721), float64(0.72031), float64(0.73286), float64(0.74632), float64(0.76432), float64(0.78448), float64(0.80421), float64(0.82199), float64(0.84363), float64(0.87423), float64(0.90617), float64(0.93827), float64(1.0)}
var N_S_Gothic int32 = 51
var S_Gothic = [51]float64{float64(0.0), float64(0.005), float64(0.0174), float64(0.03098), float64(0.04272), float64(0.055), float64(0.0698), float64(0.0862), float64(0.10461), float64(0.12463), float64(0.145), float64(0.16309), float64(0.18118), float64(0.2), float64(0.22181), float64(0.24487), float64(0.26888), float64(0.2938), float64(0.31901), float64(0.34389), float64(0.36564), float64(0.38612), float64(0.4072), float64(0.43), float64(0.45868), float64(0.48895), float64(0.52), float64(0.55032), float64(0.5804), float64(0.61), float64(0.63762), float64(0.66505), float64(0.6929), float64(0.72342), float64(0.75467), float64(0.785), float64(0.81165), float64(0.83654), float64(0.86), float64(0.88253), float64(0.90414), float64(0.925), float64(0.94486), float64(0.96475), float64(0.98567), float64(1.00833), float64(1.03), float64(1.0536), float64(1.065), float64(1.055), float64(1.0)}
var N_W_Gothic int32 = 21
var W_Gothic = [21]float64{float64(0.0), float64(0.286), float64(0.643), float64(0.762), float64(0.833), float64(0.905), float64(0.952), float64(0.976), float64(0.976), float64(1.0), float64(1.0), float64(0.976), float64(0.976), float64(0.952), float64(0.905), float64(0.833), float64(0.762), float64(0.667), float64(0.524), float64(0.357), float64(0.0)}
var N_Y_Catenary int32 = 51
var Y_Catenary = [51]float64{float64(0.0), float64(0.02974), float64(0.06439), float64(0.08433), float64(0.10549), float64(0.12064), float64(0.13952), float64(0.1556), float64(0.17032), float64(0.18512), float64(0.20057), float64(0.21995), float64(0.24011), float64(0.25892), float64(0.27595), float64(0.29214), float64(0.30802), float64(0.32372), float64(0.33894), float64(0.35315), float64(0.36557), float64(0.37833), float64(0.3923), float64(0.4097), float64(0.42982), float64(0.45), float64(0.46769), float64(0.48431), float64(0.5), float64(0.51466), float64(0.52886), float64(0.54292), float64(0.55729), float64(0.57223), float64(0.5878), float64(0.60428), float64(0.62197), float64(0.64047), float64(0.6598), float64(0.67976), float64(0.7), float64(0.71731), float64(0.73769), float64(0.76651), float64(0.8), float64(0.8209), float64(0.84311), float64(0.87978), float64(0.91576), float64(0.95), float64(1.0)}
var N_S_Catenary int32 = 51
var S_Catenary = [51]float64{float64(0.0), float64(0.00605), float64(0.01455), float64(0.0254), float64(0.03863), float64(0.0543), float64(0.07127), float64(0.08778), float64(0.10372), float64(0.12081), float64(0.14082), float64(0.16375), float64(0.18779), float64(0.21157), float64(0.23478), float64(0.25818), float64(0.28244), float64(0.30741), float64(0.33204), float64(0.35505), float64(0.37465), float64(0.39404), float64(0.41426), float64(0.43804), float64(0.46531), float64(0.49357), float64(0.52187), float64(0.54925), float64(0.57647), float64(0.60321), float64(0.62964), float64(0.65639), float64(0.68472), float64(0.71425), float64(0.74303), float64(0.76827), float64(0.79168), float64(0.815), float64(0.84094), float64(0.86707), float64(0.89213), float64(0.91607), float64(0.94), float64(0.96604), float64(0.99), float64(1.00714), float64(1.02158), float64(1.03814), float64(1.05), float64(1.05), float64(1.0)}
var N_W_Catenary int32 = 21
var W_Catenary = [21]float64{float64(0.0), float64(0.6667), float64(0.8222), float64(0.9111), float64(0.9778), float64(1.0000), float64(1.0000), float64(0.9889), float64(0.9778), float64(0.9556), float64(0.9333), float64(0.8889), float64(0.8444), float64(0.8000), float64(0.7556), float64(0.7000), float64(0.6333), float64(0.5556), float64(0.4444), float64(0.3333), float64(0.0)}
var N_Y_SemiEllip int32 = 51
var Y_SemiEllip = [51]float64{float64(0.0), float64(0.03075), float64(0.05137), float64(0.07032), float64(0.09), float64(0.11323), float64(0.13037), float64(0.14519), float64(0.15968), float64(0.18459), float64(0.19531), float64(0.21354), float64(0.22694), float64(0.23947), float64(0.25296), float64(0.265), float64(0.27784), float64(0.29212), float64(0.3097), float64(0.32982), float64(0.35), float64(0.36738), float64(0.3839), float64(0.4), float64(0.41667), float64(0.43333), float64(0.45), float64(0.46697), float64(0.48372), float64(0.5), float64(0.51374), float64(0.52747), float64(0.54209), float64(0.5595), float64(0.57941), float64(0.6), float64(0.62), float64(0.64), float64(0.66), float64(0.68), float64(0.7), float64(0.71843), float64(0.73865), float64(0.76365), float64(0.7926), float64(0.82088), float64(0.85), float64(0.88341), float64(0.90998), float64(0.93871), float64(1.0)}
var N_S_SemiEllip int32 = 51
var S_SemiEllip = [51]float64{float64(0.0), float64(0.00438), float64(0.01227), float64(0.02312), float64(0.03638), float64(0.05145), float64(0.06783), float64(0.085), float64(0.10093), float64(0.11752), float64(0.1353), float64(0.15626), float64(0.17917), float64(0.20296), float64(0.22654), float64(0.24962), float64(0.27269), float64(0.29568), float64(0.31848), float64(0.34152), float64(0.365), float64(0.38941), float64(0.41442), float64(0.44), float64(0.46636), float64(0.49309), float64(0.52), float64(0.54628), float64(0.57285), float64(0.6), float64(0.62949), float64(0.65877), float64(0.68624), float64(0.71017), float64(0.73304), float64(0.75578), float64(0.77925), float64(0.80368), float64(0.83114), float64(0.8595), float64(0.88592), float64(0.90848), float64(0.93), float64(0.95292), float64(0.97481), float64(0.99374), float64(1.01084), float64(1.02858), float64(1.04543), float64(1.05), float64(1.0)}
var N_W_SemiEllip int32 = 21
var W_SemiEllip = [21]float64{float64(0.0), float64(0.7000), float64(0.9800), float64(1.0000), float64(1.0000), float64(1.0000), float64(0.9900), float64(0.9800), float64(0.9600), float64(0.9400), float64(0.9100), float64(0.8800), float64(0.8400), float64(0.8000), float64(0.7500), float64(0.7000), float64(0.6400), float64(0.5600), float64(0.4600), float64(0.3400), float64(0.0)}
var N_A_Baskethandle int32 = 26
var A_Baskethandle = [26]float64{float64(0.0000), float64(0.0173), float64(0.0457), float64(0.0828), float64(0.1271), float64(0.1765), float64(0.2270), float64(0.2775), float64(0.3280), float64(0.3780), float64(0.4270), float64(0.4765), float64(0.5260), float64(0.5740), float64(0.6220), float64(0.6690), float64(0.7160), float64(0.7610), float64(0.8030), float64(0.8390), float64(0.8770), float64(0.9110), float64(0.9410), float64(0.9680), float64(0.9880), float64(1.000)}
var N_R_Baskethandle int32 = 26
var R_Baskethandle = [26]float64{float64(0.0100), float64(0.0952), float64(0.1890), float64(0.2730), float64(0.3690), float64(0.4630), float64(0.5600), float64(0.6530), float64(0.7430), float64(0.8220), float64(0.8830), float64(0.9490), float64(0.9990), float64(1.055), float64(1.095), float64(1.141), float64(1.161), float64(1.188), float64(1.206), float64(1.206), float64(1.206), float64(1.205), float64(1.196), float64(1.168), float64(1.127), float64(1.000)}
var N_Y_BasketHandle int32 = 51
var Y_BasketHandle = [51]float64{float64(0.0), float64(0.04112), float64(0.0738), float64(0.1), float64(0.12236), float64(0.14141), float64(0.15857), float64(0.17462), float64(0.18946), float64(0.20315), float64(0.21557), float64(0.22833), float64(0.2423), float64(0.25945), float64(0.27936), float64(0.3), float64(0.3204), float64(0.34034), float64(0.35892), float64(0.37595), float64(0.39214), float64(0.40802), float64(0.42372), float64(0.43894), float64(0.45315), float64(0.46557), float64(0.47833), float64(0.4923), float64(0.50945), float64(0.52936), float64(0.55), float64(0.57), float64(0.59), float64(0.61023), float64(0.63045), float64(0.65), float64(0.66756), float64(0.68413), float64(0.7), float64(0.71481), float64(0.72984), float64(0.74579), float64(0.76417), float64(0.78422), float64(0.80477), float64(0.82532), float64(0.85), float64(0.88277), float64(0.915), float64(0.95), float64(1.0)}
var N_S_BasketHandle int32 = 51
var S_BasketHandle = [51]float64{float64(0.0), float64(0.00758), float64(0.01812), float64(0.03), float64(0.03966), float64(0.04957), float64(0.0623), float64(0.07849), float64(0.09618), float64(0.11416), float64(0.13094), float64(0.14808), float64(0.16583), float64(0.18381), float64(0.20294), float64(0.225), float64(0.2547), float64(0.28532), float64(0.31006), float64(0.32804), float64(0.34555), float64(0.36944), float64(0.40032), float64(0.43203), float64(0.46004), float64(0.47849), float64(0.49591), float64(0.51454), float64(0.5381), float64(0.56711), float64(0.6), float64(0.64092), float64(0.68136), float64(0.71259), float64(0.73438), float64(0.755), float64(0.78625), float64(0.8188), float64(0.85), float64(0.8679), float64(0.88483), float64(0.90431), float64(0.9369), float64(0.97388), float64(1.00747), float64(1.033), float64(1.05), float64(1.05464), float64(1.06078), float64(1.055), float64(1.0)}
var N_W_BasketHandle int32 = 26
var W_BasketHandle = [26]float64{float64(0.0), float64(0.49), float64(0.667), float64(0.82), float64(0.93), float64(1.00), float64(1.00), float64(1.00), float64(0.997), float64(0.994), float64(0.988), float64(0.982), float64(0.967), float64(0.948), float64(0.928), float64(0.904), float64(0.874), float64(0.842), float64(0.798), float64(0.75), float64(0.697), float64(0.637), float64(0.567), float64(0.467), float64(0.342), float64(0.0)}
var N_Y_SemiCirc int32 = 51
var Y_SemiCirc = [51]float64{float64(0.0), float64(0.04102), float64(0.07407), float64(0.1), float64(0.11769), float64(0.13037), float64(0.14036), float64(0.15), float64(0.16546), float64(0.18213), float64(0.2), float64(0.22018), float64(0.2403), float64(0.25788), float64(0.27216), float64(0.285), float64(0.29704), float64(0.30892), float64(0.32128), float64(0.33476), float64(0.35), float64(0.36927), float64(0.38963), float64(0.41023), float64(0.43045), float64(0.45), float64(0.46769), float64(0.48431), float64(0.5), float64(0.51443), float64(0.52851), float64(0.54271), float64(0.55774), float64(0.57388), float64(0.59101), float64(0.60989), float64(0.63005), float64(0.65), float64(0.66682), float64(0.68318), float64(0.7), float64(0.71675), float64(0.73744), float64(0.76651), float64(0.8), float64(0.8209), float64(0.84311), float64(0.87978), float64(0.91576), float64(0.95), float64(1.0)}
var N_S_SemiCirc int32 = 51
var S_SemiCirc = [51]float64{float64(0.0), float64(0.00757), float64(0.01815), float64(0.03), float64(0.0358), float64(0.04037), float64(0.04601), float64(0.055), float64(0.07475), float64(0.09834), float64(0.125), float64(0.1557), float64(0.18588), float64(0.20883), float64(0.223), float64(0.23472), float64(0.24667), float64(0.26758), float64(0.29346), float64(0.32124), float64(0.35), float64(0.3772), float64(0.4054), float64(0.43541), float64(0.46722), float64(0.5), float64(0.53532), float64(0.56935), float64(0.6), float64(0.61544), float64(0.62811), float64(0.6417), float64(0.66598), float64(0.7001), float64(0.73413), float64(0.76068), float64(0.78027), float64(0.8), float64(0.82891), float64(0.85964), float64(0.89), float64(0.9127), float64(0.93664), float64(0.96677), float64(1.0), float64(1.02661), float64(1.04631), float64(1.05726), float64(1.06637), float64(1.06), float64(1.0)}
var N_W_SemiCirc int32 = 21
var W_SemiCirc = [21]float64{float64(0.0), float64(0.5488), float64(0.8537), float64(1.0000), float64(1.0000), float64(0.9939), float64(0.9878), float64(0.9756), float64(0.9634), float64(0.9451), float64(0.9207), float64(0.8902), float64(0.8537), float64(0.8171), float64(0.7683), float64(0.7073), float64(0.6463), float64(0.5732), float64(0.4756), float64(0.3354), float64(0.0)}
var NumCodesEllipse int32 = 23
var MinorAxis_Ellipse = [23]float64{float64(14), float64(19), float64(22), float64(24), float64(27), float64(29), float64(32), float64(34), float64(38), float64(43), float64(48), float64(53), float64(58), float64(63), float64(68), float64(72), float64(77), float64(82), float64(87), float64(92), float64(97), float64(106), float64(116)}
var MajorAxis_Ellipse = [23]float64{float64(23), float64(30), float64(34), float64(38), float64(42), float64(45), float64(49), float64(53), float64(60), float64(68), float64(76), float64(83), float64(91), float64(98), float64(106), float64(113), float64(121), float64(128), float64(136), float64(143), float64(151), float64(166), float64(180)}
var Afull_Ellipse = [23]float64{float64(1.80), float64(3.30), float64(4.10), float64(5.10), float64(6.30), float64(7.40), float64(8.80), float64(10.20), float64(12.90), float64(16.60), float64(20.50), float64(24.80), float64(29.50), float64(34.60), float64(40.10), float64(46.10), float64(52.40), float64(59.20), float64(66.40), float64(74.00), float64(82.00), float64(99.20), float64(118.60)}
var Rfull_Ellipse = [23]float64{float64(0.367), float64(0.490), float64(0.546), float64(0.613), float64(0.686), float64(0.736), float64(0.812), float64(0.875), float64(0.969), float64(1.106), float64(1.229), float64(1.352), float64(1.475), float64(1.598), float64(1.721), float64(1.845), float64(1.967), float64(2.091), float64(2.215), float64(2.340), float64(2.461), float64(2.707), float64(2.968)}
var N_A_HorizEllipse int32 = 26
var A_HorizEllipse = [26]float64{float64(0.0000), float64(0.0150), float64(0.0400), float64(0.0650), float64(0.0950), float64(0.1300), float64(0.1650), float64(0.2050), float64(0.2500), float64(0.3000), float64(0.3550), float64(0.4150), float64(0.4800), float64(0.5200), float64(0.5850), float64(0.6450), float64(0.7000), float64(0.7500), float64(0.7950), float64(0.8350), float64(0.8700), float64(0.9050), float64(0.9350), float64(0.9600), float64(0.9850), float64(1.000)}
var N_R_HorizEllipse int32 = 26
var R_HorizEllipse = [26]float64{float64(0.0100), float64(0.0764), float64(0.1726), float64(0.2389), float64(0.3274), float64(0.4191), float64(0.5120), float64(0.5983), float64(0.6757), float64(0.7630), float64(0.8326), float64(0.9114), float64(0.9702), float64(1.030), float64(1.091), float64(1.146), float64(1.185), float64(1.225), float64(1.257), float64(1.274), float64(1.290), float64(1.282), float64(1.274), float64(1.257), float64(1.185), float64(1.000)}
var N_W_HorizEllipse int32 = 26
var W_HorizEllipse = [26]float64{float64(0.0), float64(0.3919), float64(0.5426), float64(0.6499), float64(0.7332), float64(0.8000), float64(0.8542), float64(0.8980), float64(0.9330), float64(0.9600), float64(0.9798), float64(0.9928), float64(0.9992), float64(0.9992), float64(0.9928), float64(0.9798), float64(0.9600), float64(0.9330), float64(0.8980), float64(0.8542), float64(0.8000), float64(0.7332), float64(0.6499), float64(0.5426), float64(0.3919), float64(0.0)}
var N_A_VertEllipse int32 = 26
var A_VertEllipse = [26]float64{float64(0.0000), float64(0.0100), float64(0.0400), float64(0.0700), float64(0.1000), float64(0.1400), float64(0.1850), float64(0.2300), float64(0.2800), float64(0.3300), float64(0.3800), float64(0.4300), float64(0.4800), float64(0.5200), float64(0.5700), float64(0.6200), float64(0.6700), float64(0.7200), float64(0.7700), float64(0.8150), float64(0.8600), float64(0.9000), float64(0.9300), float64(0.9600), float64(0.9900), float64(1.000)}
var N_R_VertEllipse int32 = 26
var R_VertEllipse = [26]float64{float64(0.0100), float64(0.1250), float64(0.2436), float64(0.3536), float64(0.4474), float64(0.5484), float64(0.6366), float64(0.7155), float64(0.7768), float64(0.8396), float64(0.8969), float64(0.9480), float64(0.9925), float64(1.023), float64(1.053), float64(1.084), float64(1.107), float64(1.130), float64(1.154), float64(1.170), float64(1.177), float64(1.177), float64(1.170), float64(1.162), float64(1.122), float64(1.000)}
var N_W_VertEllipse int32 = 26
var W_VertEllipse = [26]float64{float64(0.0), float64(0.3919), float64(0.5426), float64(0.6499), float64(0.7332), float64(0.8000), float64(0.8542), float64(0.8980), float64(0.9330), float64(0.9600), float64(0.9798), float64(0.9928), float64(0.9992), float64(0.9992), float64(0.9928), float64(0.9798), float64(0.9600), float64(0.9330), float64(0.8980), float64(0.8542), float64(0.8000), float64(0.7332), float64(0.6499), float64(0.5426), float64(0.3919), float64(0.0)}
var NumCodesArch int32 = 102
var Yfull_Arch = [102]float64{float64(11), float64(13.5), float64(15.5), float64(18), float64(22.5), float64(26.625), float64(31.3125), float64(36), float64(40), float64(45), float64(54), float64(62), float64(72), float64(77.5), float64(87.125), float64(96.875), float64(106.5), float64(13), float64(15), float64(18), float64(20), float64(24), float64(29), float64(33), float64(38), float64(43), float64(47), float64(52), float64(57), float64(31), float64(36), float64(41), float64(46), float64(51), float64(55), float64(59), float64(63), float64(67), float64(71), float64(75), float64(79), float64(83), float64(87), float64(91), float64(55), float64(57), float64(59), float64(61), float64(63), float64(65), float64(67), float64(69), float64(71), float64(73), float64(75), float64(77), float64(79), float64(81), float64(83), float64(85), float64(87), float64(89), float64(91), float64(93), float64(95), float64(97), float64(100), float64(101), float64(103), float64(105), float64(107), float64(109), float64(111), float64(113), float64(115), float64(118), float64(119), float64(121), float64(112), float64(114), float64(116), float64(118), float64(120), float64(122), float64(124), float64(126), float64(128), float64(130), float64(132), float64(134), float64(136), float64(138), float64(140), float64(142), float64(144), float64(146), float64(148), float64(150), float64(152), float64(154), float64(156), float64(158)}
var Wmax_Arch = [102]float64{float64(18), float64(22), float64(26), float64(28.5), float64(36.25), float64(43.75), float64(51.125), float64(58.5), float64(65), float64(73), float64(88), float64(102), float64(115), float64(122), float64(138), float64(154), float64(168.75), float64(17), float64(21), float64(24), float64(28), float64(35), float64(42), float64(49), float64(57), float64(64), float64(71), float64(77), float64(83), float64(40), float64(46), float64(53), float64(60), float64(66), float64(73), float64(81), float64(87), float64(95), float64(103), float64(112), float64(117), float64(128), float64(137), float64(142), float64(73), float64(76), float64(81), float64(84), float64(87), float64(92), float64(95), float64(98), float64(103), float64(106), float64(112), float64(114), float64(117), float64(123), float64(128), float64(131), float64(137), float64(139), float64(142), float64(148), float64(150), float64(152), float64(154), float64(161), float64(167), float64(169), float64(171), float64(178), float64(184), float64(186), float64(188), float64(190), float64(197), float64(199), float64(159), float64(162), float64(168), float64(170), float64(173), float64(179), float64(184), float64(187), float64(190), float64(195), float64(198), float64(204), float64(206), float64(209), float64(215), float64(217), float64(223), float64(225), float64(231), float64(234), float64(236), float64(239), float64(245), float64(247)}
var Afull_Arch = [102]float64{float64(1.1), float64(1.65), float64(2.2), float64(2.8), float64(4.4), float64(6.4), float64(8.8), float64(11.4), float64(14.3), float64(17.7), float64(25.6), float64(34.6), float64(44.5), float64(51.7), float64(66), float64(81.8), float64(99.1), float64(1.1), float64(1.6), float64(2.2), float64(2.9), float64(4.5), float64(6.5), float64(8.9), float64(11.6), float64(14.7), float64(18.1), float64(21.9), float64(26), float64(7), float64(9.4), float64(12.3), float64(15.6), float64(19.3), float64(23.2), float64(27.4), float64(32.1), float64(37), float64(42.4), float64(48), float64(54.2), float64(60.5), float64(67.4), float64(74.5), float64(22), float64(24), float64(26), float64(28), float64(31), float64(33), float64(35), float64(38), float64(40), float64(43), float64(46), float64(49), float64(52), float64(55), float64(58), float64(61), float64(64), float64(67), float64(71), float64(74), float64(78), float64(81), float64(85), float64(89), float64(93), float64(97), float64(101), float64(105), float64(109), float64(113), float64(118), float64(122), float64(126), float64(131), float64(97), float64(102), float64(105), float64(109), float64(114), float64(118), float64(123), float64(127), float64(132), float64(137), float64(142), float64(146), float64(151), float64(157), float64(161), float64(167), float64(172), float64(177), float64(182), float64(188), float64(194), float64(200), float64(205), float64(211)}
var Rfull_Arch = [102]float64{float64(0.25), float64(0.3), float64(0.36), float64(0.45), float64(0.56), float64(0.68), float64(0.8), float64(0.9), float64(1.01), float64(1.13), float64(1.35), float64(1.57), float64(1.77), float64(1.92), float64(2.17), float64(2.42), float64(2.65), float64(0.324), float64(0.374), float64(0.449), float64(0.499), float64(0.598), float64(0.723), float64(0.823), float64(0.947), float64(1.072), float64(1.171), float64(1.296), float64(1.421), float64(0.773), float64(0.773), float64(1.022), float64(1.147), float64(1.271), float64(1.371), float64(1.471), float64(1.570), float64(1.670), float64(1.770), float64(1.869), float64(1.969), float64(2.069), float64(2.168), float64(2.268), float64(1.371), float64(1.421), float64(1.471), float64(1.520), float64(1.570), float64(1.620), float64(1.670), float64(1.720), float64(1.770), float64(1.820), float64(1.869), float64(1.919), float64(1.969), float64(2.019), float64(2.069), float64(2.119), float64(2.168), float64(2.218), float64(2.268), float64(2.318), float64(2.368), float64(2.418), float64(2.493), float64(2.517), float64(2.567), float64(2.617), float64(2.667), float64(2.717), float64(2.767), float64(2.817), float64(2.866), float64(2.941), float64(2.966), float64(3.016), float64(2.792), float64(2.841), float64(2.891), float64(2.941), float64(2.991), float64(3.041), float64(3.091), float64(3.141), float64(3.190), float64(3.240), float64(3.290), float64(3.340), float64(3.390), float64(3.440), float64(3.490), float64(3.539), float64(3.589), float64(3.639), float64(3.689), float64(3.739), float64(3.789), float64(3.838), float64(3.888), float64(3.938)}
var N_A_Arch int32 = 26
var A_Arch = [26]float64{float64(0.0000), float64(0.0200), float64(0.0600), float64(0.1000), float64(0.1400), float64(0.1900), float64(0.2400), float64(0.2900), float64(0.3400), float64(0.3900), float64(0.4400), float64(0.4900), float64(0.5400), float64(0.5900), float64(0.6400), float64(0.6900), float64(0.7350), float64(0.7800), float64(0.8200), float64(0.8600), float64(0.8950), float64(0.9300), float64(0.9600), float64(0.9850), float64(0.9950), float64(1.000)}
var N_R_Arch int32 = 26
var R_Arch = [26]float64{float64(0.0100), float64(0.0983), float64(0.1965), float64(0.2948), float64(0.3940), float64(0.4962), float64(0.5911), float64(0.6796), float64(0.7615), float64(0.8364), float64(0.9044), float64(0.9640), float64(1.018), float64(1.065), float64(1.106), float64(1.142), float64(1.170), float64(1.192), float64(1.208), float64(1.217), float64(1.220), float64(1.213), float64(1.196), float64(1.168), float64(1.112), float64(1.000)}
var N_W_Arch int32 = 26
var W_Arch = [26]float64{float64(0.0), float64(0.6272), float64(0.8521), float64(0.9243), float64(0.9645), float64(0.9846), float64(0.9964), float64(0.9988), float64(0.9917), float64(0.9811), float64(0.9680), float64(0.9515), float64(0.9314), float64(0.9101), float64(0.8864), float64(0.8592), float64(0.8284), float64(0.7917), float64(0.7527), float64(0.7065), float64(0.6544), float64(0.5953), float64(0.5231), float64(0.4355), float64(0.3195), float64(0.0)}
var Amax = [26]float64{float64(1.0), float64(0.9756), float64(0.9756), float64(0.97), float64(1.0), float64(1.0), float64(1.0), float64(1.0), float64(1.0), float64(0.98), float64(0.98), float64(0.96), float64(0.96), float64(0.96), float64(0.92), float64(0.96), float64(0.96), float64(0.96), float64(0.98), float64(0.98), float64(0.96), float64(0.96), float64(1.0), float64(0.96), float64(0.9756), float64(1.0)}

// xsect.c:204
func xsect_isOpen(type_2 int32) int32 {
	if Amax[type_2] >= float64(1.0) {
		return 1
	}
	return 0
}

// xsect.c:216
func xsect_setParams(xsect_2 *TXsect, type_2 int32, p []float64, ucf float64) int32 {
	var (
		index int32
		aMax  float64
		theta float64
	)
	_, _, _ = index, aMax, theta
	index = 0
	aMax = 0
	theta = 0
	if (type_2 != DUMMY) && (p[0] <= float64(0.0)) {
		return 0
	}
	xsect_2.type_ = type_2
	switch xsect_2.type_ {
	case DUMMY:
		xsect_2.yFull = float64(1e-06)
		xsect_2.wMax = float64(1e-06)
		xsect_2.aFull = float64(1e-06)
		xsect_2.rFull = float64(1e-06)
		xsect_2.sFull = float64(1e-06)
		xsect_2.sMax = float64(1e-06)
		break
	case CIRCULAR:
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.wMax = xsect_2.yFull
		xsect_2.aFull = float64(float64(float64(float64(3.141592654)/float64(4.0))*xsect_2.yFull) * xsect_2.yFull)
		xsect_2.rFull = float64(float64(0.2500) * xsect_2.yFull)
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = float64(float64(1.08) * xsect_2.sFull)
		xsect_2.ywMax = float64(float64(0.5) * xsect_2.yFull)
		break
	case FORCE_MAIN:
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.wMax = xsect_2.yFull
		xsect_2.aFull = float64(float64(float64(float64(3.141592654)/float64(4.0))*xsect_2.yFull) * xsect_2.yFull)
		xsect_2.rFull = float64(float64(0.2500) * xsect_2.yFull)
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(0.63)))
		xsect_2.sMax = float64(float64(1.06949) * xsect_2.sFull)
		xsect_2.ywMax = float64(float64(0.5) * xsect_2.yFull)
		xsect_2.rBot = p[1]
		break
	case FILLED_CIRCULAR:
		if p[1] >= p[0] {
			return 0
		}
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.wMax = xsect_2.yFull
		xsect_2.aFull = float64(float64(float64(float64(3.141592654)/float64(4.0))*xsect_2.yFull) * xsect_2.yFull)
		xsect_2.rFull = float64(float64(0.2500) * xsect_2.yFull)
		xsect_2.yBot = float64(p[1] / ucf)
		xsect_2.aBot = circ_getAofY(xsect_2, xsect_2.yBot)
		xsect_2.sBot = xsect_getWofY(xsect_2, xsect_2.yBot)
		xsect_2.rBot = float64(xsect_2.aBot / (float64(xsect_2.rFull * lookup(float64(xsect_2.yBot/xsect_2.yFull), R_Circ[:], N_R_Circ))))
		xsect_2.aFull = float64(xsect_2.aFull - xsect_2.aBot)
		xsect_2.rFull = float64(xsect_2.aFull / (float64(float64(float64(float64(3.141592654)*xsect_2.yFull)-xsect_2.rBot) + xsect_2.sBot)))
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = float64(float64(1.08) * xsect_2.sFull)
		xsect_2.yFull = float64(xsect_2.yFull - xsect_2.yBot)
		xsect_2.ywMax = float64(float64(0.5) * xsect_2.yFull)
		break
	case EGGSHAPED:
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.aFull = float64(float64(float64(0.5105)*xsect_2.yFull) * xsect_2.yFull)
		xsect_2.rFull = float64(float64(0.1931) * xsect_2.yFull)
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = float64(float64(1.065) * xsect_2.sFull)
		xsect_2.wMax = float64(float64(float64(2.0)/float64(3.0)) * xsect_2.yFull)
		xsect_2.ywMax = float64(float64(0.64) * xsect_2.yFull)
		break
	case HORSESHOE:
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.aFull = float64(float64(float64(0.8293)*xsect_2.yFull) * xsect_2.yFull)
		xsect_2.rFull = float64(float64(0.2538) * xsect_2.yFull)
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = float64(float64(1.077) * xsect_2.sFull)
		xsect_2.wMax = float64(float64(1.0) * xsect_2.yFull)
		xsect_2.ywMax = float64(float64(0.5) * xsect_2.yFull)
		break
	case GOTHIC:
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.aFull = float64(float64(float64(0.6554)*xsect_2.yFull) * xsect_2.yFull)
		xsect_2.rFull = float64(float64(0.2269) * xsect_2.yFull)
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = float64(float64(1.065) * xsect_2.sFull)
		xsect_2.wMax = float64(float64(0.84) * xsect_2.yFull)
		xsect_2.ywMax = float64(float64(0.45) * xsect_2.yFull)
		break
	case CATENARY:
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.aFull = float64(float64(float64(0.70277)*xsect_2.yFull) * xsect_2.yFull)
		xsect_2.rFull = float64(float64(0.23172) * xsect_2.yFull)
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = float64(float64(1.05) * xsect_2.sFull)
		xsect_2.wMax = float64(float64(0.9) * xsect_2.yFull)
		xsect_2.ywMax = float64(float64(0.25) * xsect_2.yFull)
		break
	case SEMIELLIPTICAL:
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.aFull = float64(float64(float64(0.785)*xsect_2.yFull) * xsect_2.yFull)
		xsect_2.rFull = float64(float64(0.242) * xsect_2.yFull)
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = float64(float64(1.045) * xsect_2.sFull)
		xsect_2.wMax = float64(float64(1.0) * xsect_2.yFull)
		xsect_2.ywMax = float64(float64(0.15) * xsect_2.yFull)
		break
	case BASKETHANDLE:
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.aFull = float64(float64(float64(0.7862)*xsect_2.yFull) * xsect_2.yFull)
		xsect_2.rFull = float64(float64(0.2464) * xsect_2.yFull)
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = float64(float64(1.06078) * xsect_2.sFull)
		xsect_2.wMax = float64(float64(0.944) * xsect_2.yFull)
		xsect_2.ywMax = float64(float64(0.2) * xsect_2.yFull)
		break
	case SEMICIRCULAR:
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.aFull = float64(float64(float64(1.2697)*xsect_2.yFull) * xsect_2.yFull)
		xsect_2.rFull = float64(float64(0.2946) * xsect_2.yFull)
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = float64(float64(1.06637) * xsect_2.sFull)
		xsect_2.wMax = float64(float64(1.64) * xsect_2.yFull)
		xsect_2.ywMax = float64(float64(0.15) * xsect_2.yFull)
		break
	case RECT_CLOSED:
		if p[1] <= float64(0.0) {
			return 0
		}
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.wMax = float64(p[1] / ucf)
		xsect_2.aFull = float64(xsect_2.yFull * xsect_2.wMax)
		xsect_2.rFull = float64(xsect_2.aFull / (float64(float64(2.0) * (float64(xsect_2.yFull + xsect_2.wMax)))))
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		aMax = float64(float64(0.97) * xsect_2.aFull)
		xsect_2.sMax = float64(aMax * M.Pow(rect_closed_getRofA(xsect_2, aMax), float64(float64(2.0)/float64(3.0))))
		xsect_2.ywMax = xsect_2.yFull
		break
	case RECT_OPEN:
		if p[1] <= float64(0.0) {
			return 0
		}
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.wMax = float64(p[1] / ucf)
		if (p[2] < float64(0.0)) || (p[2] > float64(2.0)) {
			return 0
		}
		xsect_2.sBot = p[2]
		xsect_2.aFull = float64(xsect_2.yFull * xsect_2.wMax)
		xsect_2.rFull = float64(xsect_2.aFull / (float64(float64((float64(float64(2.0)-xsect_2.sBot))*xsect_2.yFull) + xsect_2.wMax)))
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = xsect_2.sFull
		xsect_2.ywMax = xsect_2.yFull
		break
	case RECT_TRIANG:
		if (p[1] <= float64(0.0)) || (p[2] <= float64(0.0)) {
			return 0
		}
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.wMax = float64(p[1] / ucf)
		xsect_2.yBot = float64(p[2] / ucf)
		xsect_2.ywMax = xsect_2.yFull
		xsect_2.aBot = float64(float64(xsect_2.yBot*xsect_2.wMax) / float64(2.0))
		xsect_2.sBot = float64(float64(xsect_2.wMax/xsect_2.yBot) / float64(2.0))
		xsect_2.rBot = M.Sqrt(float64(float64(1.0) + float64(xsect_2.sBot*xsect_2.sBot)))
		xsect_2.aFull = float64(xsect_2.wMax * (float64(xsect_2.yFull - float64(xsect_2.yBot/float64(2.0)))))
		xsect_2.rFull = float64(xsect_2.aFull / (float64(float64(float64(float64(float64(2.0)*xsect_2.yBot)*xsect_2.rBot)+float64(float64(2.0)*(float64(xsect_2.yFull-xsect_2.yBot)))) + xsect_2.wMax)))
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		aMax = float64(float64(0.98) * xsect_2.aFull)
		xsect_2.sMax = float64(aMax * M.Pow(rect_triang_getRofA(xsect_2, aMax), float64(float64(2.0)/float64(3.0))))
		break
	case RECT_ROUND:
		if p[1] <= float64(0.0) {
			return 0
		}
		if p[2] < float64(p[1]/float64(2.0)) {
			p[2] = float64(p[1] / float64(2.0))
		}
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.wMax = float64(p[1] / ucf)
		xsect_2.rBot = float64(p[2] / ucf)
		theta = float64(float64(2.0) * M.Asin(float64(float64(xsect_2.wMax/float64(2.0))/xsect_2.rBot)))
		xsect_2.aBot = float64(float64(float64(xsect_2.rBot*xsect_2.rBot)/float64(2.0)) * (float64(theta - M.Sin(theta))))
		xsect_2.sBot = float64(float64(float64(float64(3.141592654)*xsect_2.rBot)*xsect_2.rBot) * M.Pow(float64(xsect_2.rBot/float64(2.0)), float64(float64(2.0)/float64(3.0))))
		xsect_2.yBot = float64(xsect_2.rBot * (float64(float64(1.0) - M.Cos(float64(theta/float64(2.0))))))
		if xsect_2.yBot > xsect_2.yFull {
			return 0
		}
		xsect_2.ywMax = xsect_2.yFull
		xsect_2.aFull = float64(float64(xsect_2.wMax*(float64(xsect_2.yFull-xsect_2.yBot))) + xsect_2.aBot)
		xsect_2.rFull = float64(xsect_2.aFull / (float64(float64(float64(xsect_2.rBot*theta)+float64(float64(2.0)*(float64(xsect_2.yFull-xsect_2.yBot)))) + xsect_2.wMax)))
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		aMax = float64(float64(0.98) * xsect_2.aFull)
		xsect_2.sMax = float64(aMax * M.Pow(rect_round_getRofA(xsect_2, aMax), float64(float64(2.0)/float64(3.0))))
		break
	case MOD_BASKET:
		if p[1] <= float64(0.0) {
			return 0
		}
		if p[2] < float64(p[1]/float64(2.0)) {
			p[2] = float64(p[1] / float64(2.0))
		}
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.wMax = float64(p[1] / ucf)
		xsect_2.rBot = float64(p[2] / ucf)
		theta = float64(float64(2.0) * M.Asin(float64(float64(xsect_2.wMax/float64(2.0))/xsect_2.rBot)))
		xsect_2.sBot = theta
		xsect_2.yBot = float64(xsect_2.rBot * (float64(float64(1.0) - M.Cos(float64(theta/float64(2.0))))))
		if xsect_2.yBot > xsect_2.yFull {
			return 0
		}
		xsect_2.ywMax = float64(xsect_2.yFull - xsect_2.yBot)
		xsect_2.aBot = float64(float64(float64(xsect_2.rBot*xsect_2.rBot)/float64(2.0)) * (float64(theta - M.Sin(theta))))
		xsect_2.aFull = float64(float64((float64(xsect_2.yFull-xsect_2.yBot))*xsect_2.wMax) + xsect_2.aBot)
		xsect_2.rFull = float64(xsect_2.aFull / (float64(float64(float64(xsect_2.rBot*theta)+float64(float64(2.0)*(float64(xsect_2.yFull-xsect_2.yBot)))) + xsect_2.wMax)))
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = xsect_getSofA(xsect_2, float64(Amax[MOD_BASKET]*xsect_2.aFull))
		break
	case TRAPEZOIDAL:
		if ((p[1] < float64(0.0)) || (p[2] < float64(0.0))) || (p[3] < float64(0.0)) {
			return 0
		}
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.ywMax = xsect_2.yFull
		xsect_2.yBot = float64(p[1] / ucf)
		xsect_2.sBot = float64((float64(p[2] + p[3])) / float64(2.0))
		if (xsect_2.yBot == float64(0.0)) && (xsect_2.sBot == float64(0.0)) {
			return 0
		}
		xsect_2.rBot = float64(M.Sqrt(float64(float64(1.0)+float64(p[2]*p[2]))) + M.Sqrt(float64(float64(1.0)+float64(p[3]*p[3]))))
		xsect_2.wMax = float64(xsect_2.yBot + float64(xsect_2.yFull*(float64(p[2]+p[3]))))
		xsect_2.aFull = float64((float64(xsect_2.yBot + float64(xsect_2.sBot*xsect_2.yFull))) * xsect_2.yFull)
		xsect_2.rFull = float64(xsect_2.aFull / (float64(xsect_2.yBot + float64(xsect_2.yFull*xsect_2.rBot))))
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = xsect_2.sFull
		break
	case TRIANGULAR:
		if p[1] <= float64(0.0) {
			return 0
		}
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.wMax = float64(p[1] / ucf)
		xsect_2.ywMax = xsect_2.yFull
		xsect_2.sBot = float64(float64(xsect_2.wMax/xsect_2.yFull) / float64(2.0))
		xsect_2.rBot = M.Sqrt(float64(float64(1.0) + float64(xsect_2.sBot*xsect_2.sBot)))
		xsect_2.aFull = float64(float64(xsect_2.yFull*xsect_2.yFull) * xsect_2.sBot)
		xsect_2.rFull = float64(xsect_2.aFull / (float64(float64(float64(2.0)*xsect_2.yFull) * xsect_2.rBot)))
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = xsect_2.sFull
		break
	case PARABOLIC:
		if p[1] <= float64(0.0) {
			return 0
		}
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.wMax = float64(p[1] / ucf)
		xsect_2.ywMax = xsect_2.yFull
		xsect_2.rBot = float64(float64(xsect_2.wMax/float64(2.0)) / M.Sqrt(xsect_2.yFull))
		xsect_2.aFull = float64(float64((float64(float64(2.0)/float64(3.0)))*xsect_2.yFull) * xsect_2.wMax)
		xsect_2.rFull = xsect_getRofY(xsect_2, xsect_2.yFull)
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = xsect_2.sFull
		break
	case POWERFUNC:
		if (p[1] <= float64(0.0)) || (p[2] <= float64(0.0)) {
			return 0
		}
		xsect_2.yFull = float64(p[0] / ucf)
		xsect_2.wMax = float64(p[1] / ucf)
		xsect_2.ywMax = xsect_2.yFull
		xsect_2.sBot = float64(float64(1.0) / p[2])
		xsect_2.rBot = float64(float64(xsect_2.wMax/(float64(xsect_2.sBot+float64(1)))) / M.Pow(xsect_2.yFull, xsect_2.sBot))
		xsect_2.aFull = float64(float64(xsect_2.yFull*xsect_2.wMax) / (float64(xsect_2.sBot + float64(1))))
		xsect_2.rFull = xsect_getRofY(xsect_2, xsect_2.yFull)
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = xsect_2.sFull
		break
	case HORIZ_ELLIPSE:
		if p[1] == float64(0.0) {
			p[2] = p[0]
		}
		if p[2] > float64(0.0) {
			index = (R.D2I(M.Floor(p[2])) - 1)
			if (index < 0) || (index >= NumCodesEllipse) {
				return 0
			}
			xsect_2.yFull = float64(MinorAxis_Ellipse[index] / float64(12.0))
			xsect_2.wMax = float64(MajorAxis_Ellipse[index] / float64(12.0))
			xsect_2.aFull = Afull_Ellipse[index]
			xsect_2.rFull = Rfull_Ellipse[index]
		} else {
			xsect_2.yFull = float64(p[0] / ucf)
			if p[1] < float64(0.0) {
				return 0
			}
			xsect_2.wMax = float64(p[1] / ucf)
			xsect_2.aFull = float64(float64(float64(1.2692)*xsect_2.yFull) * xsect_2.yFull)
			xsect_2.rFull = float64(float64(0.3061) * xsect_2.yFull)
		}
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = xsect_2.sFull
		xsect_2.ywMax = float64(float64(0.48) * xsect_2.yFull)
		break
	case VERT_ELLIPSE:
		if p[1] == float64(0.0) {
			p[2] = p[0]
		}
		if p[2] > float64(0.0) {
			index = (R.D2I(M.Floor(p[2])) - 1)
			if (index < 0) || (index >= NumCodesEllipse) {
				return 0
			}
			xsect_2.yFull = float64(MajorAxis_Ellipse[index] / float64(12.0))
			xsect_2.wMax = float64(MinorAxis_Ellipse[index] / float64(12.0))
			xsect_2.aFull = Afull_Ellipse[index]
			xsect_2.rFull = Rfull_Ellipse[index]
		} else {
			if p[1] < float64(0.0) {
				return 0
			}
			xsect_2.yFull = float64(p[0] / ucf)
			xsect_2.wMax = float64(p[1] / ucf)
			xsect_2.aFull = float64(float64(float64(1.2692)*xsect_2.wMax) * xsect_2.wMax)
			xsect_2.rFull = float64(float64(0.3061) * xsect_2.wMax)
		}
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = xsect_2.sFull
		xsect_2.ywMax = float64(float64(0.48) * xsect_2.yFull)
		break
	case ARCH:
		if p[1] == float64(0.0) {
			p[2] = p[0]
		}
		if p[2] > float64(0.0) {
			index = (R.D2I(M.Floor(p[2])) - 1)
			if (index < 0) || (index >= NumCodesArch) {
				return 0
			}
			xsect_2.yFull = float64(Yfull_Arch[index] / float64(12.0))
			xsect_2.wMax = float64(Wmax_Arch[index] / float64(12.0))
			xsect_2.aFull = Afull_Arch[index]
			xsect_2.rFull = Rfull_Arch[index]
		} else {
			if p[1] < float64(0.0) {
				return 0
			}
			xsect_2.yFull = float64(p[0] / ucf)
			xsect_2.wMax = float64(p[1] / ucf)
			xsect_2.aFull = float64(float64(float64(0.7879)*xsect_2.yFull) * xsect_2.wMax)
			xsect_2.rFull = float64(float64(0.2991) * xsect_2.yFull)
		}
		xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
		xsect_2.sMax = xsect_2.sFull
		xsect_2.ywMax = float64(float64(0.28) * xsect_2.yFull)
		break
	}
	return 1
}

// xsect.c:638
func xsect_setIrregXsectParams(xsect_2 *TXsect) {
	var (
		index int32
	)
	_ = index
	index = xsect_2.transect
	getTransectParams(xsect_2, &Transect[index])
}

// xsect.c:651
func xsect_setStreetXsectParams(xsect_2 *TXsect) {
	var (
		index int32
	)
	_ = index
	index = xsect_2.transect
	getTransectParams(xsect_2, &Street[index].transect)
}

// xsect.c:664
func xsect_setCustomXsectParams(xsect_2 *TXsect) {
	var (
		index int32
		yFull float64
		i     int32
		iMax  int32
		wMax  float64
		wTbl  []float64
	)
	_, _, _, _, _, _ = index, yFull, i, iMax, wMax, wTbl
	index = Curve[xsect_2.transect].refersTo
	yFull = xsect_2.yFull
	i = 0
	iMax = 0
	wMax = 0
	wTbl = Shape[index].widthTbl[:]
	xsect_2.wMax = float64(Shape[index].wMax * yFull)
	xsect_2.aFull = float64(float64(Shape[index].aFull*yFull) * yFull)
	xsect_2.rFull = float64(Shape[index].rFull * yFull)
	xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
	xsect_2.sMax = float64(float64(float64(Shape[index].sMax*yFull)*yFull) * M.Pow(yFull, float64(float64(2.0)/float64(3.0))))
	xsect_2.aBot = float64(float64(Shape[index].aMax*yFull) * yFull)
	iMax = 0
	wMax = wTbl[0]
	for i = 1; i < 51; i++ {
		if wTbl[i] < wMax {
			break
		}
		wMax = wTbl[i]
		iMax = i
	}
	xsect_2.ywMax = float64(float64(yFull*float64(iMax)) / float64((51 - 1)))
}

// xsect.c:700
func xsect_getAmax(xsect_2 *TXsect) float64 {
	if xsect_2.type_ == IRREGULAR {
		return xsect_2.aBot
	} else if xsect_2.type_ == CUSTOM {
		return xsect_2.aBot
	} else {
		return float64(Amax[xsect_2.type_] * xsect_2.aFull)
	}
	return 0
}

// xsect.c:714
func xsect_getSofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		alpha float64
		r     float64
	)
	_, _ = alpha, r
	alpha = float64(a_2 / xsect_2.aFull)
	r = 0
	switch xsect_2.type_ {
	case FORCE_MAIN, CIRCULAR:
		return circ_getSofA(xsect_2, a_2)
	case EGGSHAPED:
		return float64(xsect_2.sFull * lookup(alpha, S_Egg[:], N_S_Egg))
	case HORSESHOE:
		return float64(xsect_2.sFull * lookup(alpha, S_Horseshoe[:], N_S_Horseshoe))
	case GOTHIC:
		return float64(xsect_2.sFull * lookup(alpha, S_Gothic[:], N_S_Gothic))
	case CATENARY:
		return float64(xsect_2.sFull * lookup(alpha, S_Catenary[:], N_S_Catenary))
	case SEMIELLIPTICAL:
		return float64(xsect_2.sFull * lookup(alpha, S_SemiEllip[:], N_S_SemiEllip))
	case BASKETHANDLE:
		return float64(xsect_2.sFull * lookup(alpha, S_BasketHandle[:], N_S_BasketHandle))
	case SEMICIRCULAR:
		return float64(xsect_2.sFull * lookup(alpha, S_SemiCirc[:], N_S_SemiCirc))
	case RECT_CLOSED:
		return rect_closed_getSofA(xsect_2, a_2)
	case RECT_OPEN:
		return rect_open_getSofA(xsect_2, a_2)
	case RECT_TRIANG:
		return rect_triang_getSofA(xsect_2, a_2)
	case RECT_ROUND:
		return rect_round_getSofA(xsect_2, a_2)
	default:
		if a_2 == float64(0.0) {
			return float64(0.0)
		}
		r = xsect_getRofA(xsect_2, a_2)
		if r < float64(1e-06) {
			return float64(0.0)
		}
		return float64(a_2 * M.Pow(r, float64(float64(2.0)/float64(3.0))))
	}
	return 0
}

// xsect.c:773
func xsect_getYofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		alpha float64
	)
	_ = alpha
	alpha = float64(a_2 / xsect_2.aFull)
	switch xsect_2.type_ {
	case FORCE_MAIN, CIRCULAR:
		return circ_getYofA(xsect_2, a_2)
	case FILLED_CIRCULAR:
		return filled_circ_getYofA(xsect_2, a_2)
	case EGGSHAPED:
		return float64(xsect_2.yFull * lookup(alpha, Y_Egg[:], N_Y_Egg))
	case HORSESHOE:
		return float64(xsect_2.yFull * lookup(alpha, Y_Horseshoe[:], N_Y_Horseshoe))
	case GOTHIC:
		return float64(xsect_2.yFull * lookup(alpha, Y_Gothic[:], N_Y_Gothic))
	case CATENARY:
		return float64(xsect_2.yFull * lookup(alpha, Y_Catenary[:], N_Y_Catenary))
	case SEMIELLIPTICAL:
		return float64(xsect_2.yFull * lookup(alpha, Y_SemiEllip[:], N_Y_SemiEllip))
	case BASKETHANDLE:
		return float64(xsect_2.yFull * lookup(alpha, Y_BasketHandle[:], N_Y_BasketHandle))
	case SEMICIRCULAR:
		return float64(xsect_2.yFull * lookup(alpha, Y_SemiCirc[:], N_Y_SemiCirc))
	case HORIZ_ELLIPSE:
		return float64(xsect_2.yFull * invLookup(alpha, A_HorizEllipse[:], N_A_HorizEllipse))
	case VERT_ELLIPSE:
		return float64(xsect_2.yFull * invLookup(alpha, A_VertEllipse[:], N_A_VertEllipse))
	case IRREGULAR:
		return float64(xsect_2.yFull * invLookup(alpha, Transect[xsect_2.transect].areaTbl[:], 51))
	case CUSTOM:
		return float64(xsect_2.yFull * invLookup(alpha, Shape[Curve[xsect_2.transect].refersTo].areaTbl[:], 51))
	case STREET_XSECT:
		return float64(xsect_2.yFull * invLookup(alpha, Street[xsect_2.transect].transect.areaTbl[:], Street[xsect_2.transect].transect.nTbl))
	case ARCH:
		return float64(xsect_2.yFull * invLookup(alpha, A_Arch[:], N_A_Arch))
	case RECT_CLOSED:
		return float64(a_2 / xsect_2.wMax)
	case RECT_TRIANG:
		return rect_triang_getYofA(xsect_2, a_2)
	case RECT_ROUND:
		return rect_round_getYofA(xsect_2, a_2)
	case RECT_OPEN:
		return float64(a_2 / xsect_2.wMax)
	case MOD_BASKET:
		return mod_basket_getYofA(xsect_2, a_2)
	case TRAPEZOIDAL:
		return trapez_getYofA(xsect_2, a_2)
	case TRIANGULAR:
		return triang_getYofA(xsect_2, a_2)
	case PARABOLIC:
		return parab_getYofA(xsect_2, a_2)
	case POWERFUNC:
		return powerfunc_getYofA(xsect_2, a_2)
	default:
		return float64(0.0)
	}
	return 0
}

// xsect.c:857
func xsect_getAofY(xsect_2 *TXsect, y_2 float64) float64 {
	var (
		yNorm float64
	)
	_ = yNorm
	yNorm = float64(y_2 / xsect_2.yFull)
	if y_2 <= float64(0.0) {
		return float64(0.0)
	}
	switch xsect_2.type_ {
	case FORCE_MAIN, CIRCULAR:
		return float64(xsect_2.aFull * lookup(yNorm, A_Circ[:], N_A_Circ))
	case FILLED_CIRCULAR:
		return filled_circ_getAofY(xsect_2, y_2)
	case EGGSHAPED:
		return float64(xsect_2.aFull * lookup(yNorm, A_Egg[:], N_A_Egg))
	case HORSESHOE:
		return float64(xsect_2.aFull * lookup(yNorm, A_Horseshoe[:], N_A_Horseshoe))
	case GOTHIC:
		return float64(xsect_2.aFull * invLookup(yNorm, Y_Gothic[:], N_Y_Gothic))
	case CATENARY:
		return float64(xsect_2.aFull * invLookup(yNorm, Y_Catenary[:], N_Y_Catenary))
	case SEMIELLIPTICAL:
		return float64(xsect_2.aFull * invLookup(yNorm, Y_SemiEllip[:], N_Y_SemiEllip))
	case BASKETHANDLE:
		return float64(xsect_2.aFull * lookup(yNorm, A_Baskethandle[:], N_A_Baskethandle))
	case SEMICIRCULAR:
		return float64(xsect_2.aFull * invLookup(yNorm, Y_SemiCirc[:], N_Y_SemiCirc))
	case HORIZ_ELLIPSE:
		return float64(xsect_2.aFull * lookup(yNorm, A_HorizEllipse[:], N_A_HorizEllipse))
	case VERT_ELLIPSE:
		return float64(xsect_2.aFull * lookup(yNorm, A_VertEllipse[:], N_A_VertEllipse))
	case ARCH:
		return float64(xsect_2.aFull * lookup(yNorm, A_Arch[:], N_A_Arch))
	case IRREGULAR:
		return float64(xsect_2.aFull * lookup(yNorm, Transect[xsect_2.transect].areaTbl[:], 51))
	case CUSTOM:
		return float64(xsect_2.aFull * lookup(yNorm, Shape[Curve[xsect_2.transect].refersTo].areaTbl[:], 51))
	case STREET_XSECT:
		return float64(xsect_2.aFull * lookup(yNorm, Street[xsect_2.transect].transect.areaTbl[:], Street[xsect_2.transect].transect.nTbl))
	case RECT_CLOSED:
		return float64(y_2 * xsect_2.wMax)
	case RECT_TRIANG:
		return rect_triang_getAofY(xsect_2, y_2)
	case RECT_ROUND:
		return rect_round_getAofY(xsect_2, y_2)
	case RECT_OPEN:
		return float64(y_2 * xsect_2.wMax)
	case MOD_BASKET:
		return mod_basket_getAofY(xsect_2, y_2)
	case TRAPEZOIDAL:
		return trapez_getAofY(xsect_2, y_2)
	case TRIANGULAR:
		return triang_getAofY(xsect_2, y_2)
	case PARABOLIC:
		return parab_getAofY(xsect_2, y_2)
	case POWERFUNC:
		return powerfunc_getAofY(xsect_2, y_2)
	default:
		return float64(0.0)
	}
	return 0
}

// xsect.c:943
func xsect_getWofY(xsect_2 *TXsect, y_2 float64) float64 {
	var (
		yNorm float64
	)
	_ = yNorm
	yNorm = float64(y_2 / xsect_2.yFull)
	switch xsect_2.type_ {
	case FORCE_MAIN, CIRCULAR:
		return float64(xsect_2.wMax * lookup(yNorm, W_Circ[:], N_W_Circ))
	case FILLED_CIRCULAR:
		yNorm = float64((float64(y_2 + xsect_2.yBot)) / (float64(xsect_2.yFull + xsect_2.yBot)))
		return float64(xsect_2.wMax * lookup(yNorm, W_Circ[:], N_W_Circ))
	case EGGSHAPED:
		return float64(xsect_2.wMax * lookup(yNorm, W_Egg[:], N_W_Egg))
	case HORSESHOE:
		return float64(xsect_2.wMax * lookup(yNorm, W_Horseshoe[:], N_W_Horseshoe))
	case GOTHIC:
		return float64(xsect_2.wMax * lookup(yNorm, W_Gothic[:], N_W_Gothic))
	case CATENARY:
		return float64(xsect_2.wMax * lookup(yNorm, W_Catenary[:], N_W_Catenary))
	case SEMIELLIPTICAL:
		return float64(xsect_2.wMax * lookup(yNorm, W_SemiEllip[:], N_W_SemiEllip))
	case BASKETHANDLE:
		return float64(xsect_2.wMax * lookup(yNorm, W_BasketHandle[:], N_W_BasketHandle))
	case SEMICIRCULAR:
		return float64(xsect_2.wMax * lookup(yNorm, W_SemiCirc[:], N_W_SemiCirc))
	case HORIZ_ELLIPSE:
		return float64(xsect_2.wMax * lookup(yNorm, W_HorizEllipse[:], N_W_HorizEllipse))
	case VERT_ELLIPSE:
		return float64(xsect_2.wMax * lookup(yNorm, W_VertEllipse[:], N_W_VertEllipse))
	case ARCH:
		return float64(xsect_2.wMax * lookup(yNorm, W_Arch[:], N_W_Arch))
	case IRREGULAR:
		return float64(xsect_2.wMax * lookup(yNorm, Transect[xsect_2.transect].widthTbl[:], 51))
	case CUSTOM:
		return float64(xsect_2.wMax * lookup(yNorm, Shape[Curve[xsect_2.transect].refersTo].widthTbl[:], 51))
	case STREET_XSECT:
		return float64(xsect_2.wMax * lookup(yNorm, Street[xsect_2.transect].transect.widthTbl[:], Street[xsect_2.transect].transect.nTbl))
	case RECT_CLOSED:
		if yNorm == float64(1.0) {
			return float64(0.0)
		}
		return xsect_2.wMax
	case RECT_TRIANG:
		return rect_triang_getWofY(xsect_2, y_2)
	case RECT_ROUND:
		return rect_round_getWofY(xsect_2, y_2)
	case RECT_OPEN:
		return xsect_2.wMax
	case MOD_BASKET:
		return mod_basket_getWofY(xsect_2, y_2)
	case TRAPEZOIDAL:
		return trapez_getWofY(xsect_2, y_2)
	case TRIANGULAR:
		return triang_getWofY(xsect_2, y_2)
	case PARABOLIC:
		return parab_getWofY(xsect_2, y_2)
	case POWERFUNC:
		return powerfunc_getWofY(xsect_2, y_2)
	default:
		return float64(0.0)
	}
	return 0
}

// xsect.c:1031
func xsect_getRofY(xsect_2 *TXsect, y_2 float64) float64 {
	var (
		yNorm float64
	)
	_ = yNorm
	yNorm = float64(y_2 / xsect_2.yFull)
	switch xsect_2.type_ {
	case FORCE_MAIN, CIRCULAR:
		return float64(xsect_2.rFull * lookup(yNorm, R_Circ[:], N_R_Circ))
	case FILLED_CIRCULAR:
		if xsect_2.yBot == float64(0.0) {
			return float64(xsect_2.rFull * lookup(yNorm, R_Circ[:], N_R_Circ))
		}
		return filled_circ_getRofY(xsect_2, y_2)
	case EGGSHAPED:
		return float64(xsect_2.rFull * lookup(yNorm, R_Egg[:], N_R_Egg))
	case HORSESHOE:
		return float64(xsect_2.rFull * lookup(yNorm, R_Horseshoe[:], N_R_Horseshoe))
	case BASKETHANDLE:
		return float64(xsect_2.rFull * lookup(yNorm, R_Baskethandle[:], N_R_Baskethandle))
	case HORIZ_ELLIPSE:
		return float64(xsect_2.rFull * lookup(yNorm, R_HorizEllipse[:], N_R_HorizEllipse))
	case VERT_ELLIPSE:
		return float64(xsect_2.rFull * lookup(yNorm, R_VertEllipse[:], N_R_VertEllipse))
	case ARCH:
		return float64(xsect_2.rFull * lookup(yNorm, R_Arch[:], N_R_Arch))
	case IRREGULAR:
		return float64(xsect_2.rFull * lookup(yNorm, Transect[xsect_2.transect].hradTbl[:], 51))
	case CUSTOM:
		return float64(xsect_2.rFull * lookup(yNorm, Shape[Curve[xsect_2.transect].refersTo].hradTbl[:], 51))
	case STREET_XSECT:
		return float64(xsect_2.rFull * lookup(yNorm, Street[xsect_2.transect].transect.hradTbl[:], Street[xsect_2.transect].transect.nTbl))
	case RECT_TRIANG:
		return rect_triang_getRofY(xsect_2, y_2)
	case RECT_ROUND:
		return rect_round_getRofY(xsect_2, y_2)
	case TRAPEZOIDAL:
		return trapez_getRofY(xsect_2, y_2)
	case TRIANGULAR:
		return triang_getRofY(xsect_2, y_2)
	case PARABOLIC:
		return parab_getRofY(xsect_2, y_2)
	case POWERFUNC:
		return powerfunc_getRofY(xsect_2, y_2)
	default:
		return xsect_getRofA(xsect_2, xsect_getAofY(xsect_2, y_2))
	}
	return 0
}

// xsect.c:1100
func xsect_getRofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		cathy float64
	)
	_ = cathy
	cathy = 0
	if a_2 <= float64(0.0) {
		return float64(0.0)
	}
	switch xsect_2.type_ {
	case HORIZ_ELLIPSE, VERT_ELLIPSE, ARCH, IRREGULAR, FILLED_CIRCULAR, CUSTOM, STREET_XSECT:
		return xsect_getRofY(xsect_2, xsect_getYofA(xsect_2, a_2))
	case RECT_CLOSED:
		return rect_closed_getRofA(xsect_2, a_2)
	case RECT_OPEN:
		return float64(a_2 / (float64(xsect_2.wMax + float64(float64((float64(float64(2.0)-xsect_2.sBot))*a_2)/xsect_2.wMax))))
	case RECT_TRIANG:
		return rect_triang_getRofA(xsect_2, a_2)
	case RECT_ROUND:
		return rect_round_getRofA(xsect_2, a_2)
	case MOD_BASKET:
		return mod_basket_getRofA(xsect_2, a_2)
	case TRAPEZOIDAL:
		return trapez_getRofA(xsect_2, a_2)
	case TRIANGULAR:
		return triang_getRofA(xsect_2, a_2)
	case PARABOLIC:
		return parab_getRofA(xsect_2, a_2)
	case POWERFUNC:
		return powerfunc_getRofA(xsect_2, a_2)
	default:
		cathy = xsect_getSofA(xsect_2, a_2)
		if (cathy < float64(1e-06)) || (a_2 < float64(1e-06)) {
			return float64(0.0)
		}
		return M.Pow(float64(cathy/a_2), float64(float64(3.0)/float64(2.0)))
	}
	return 0
}

// xsect.c:1149
func xsect_getAofS(xsect_2 *TXsect, s float64) float64 {
	var (
		psi float64
	)
	_ = psi
	psi = float64(s / xsect_2.sFull)
	if s <= float64(0.0) {
		return float64(0.0)
	}
	if s > xsect_2.sMax {
		s = xsect_2.sMax
	}
	switch xsect_2.type_ {
	case DUMMY:
		return float64(0.0)
	case FORCE_MAIN, CIRCULAR:
		return circ_getAofS(xsect_2, s)
	case EGGSHAPED:
		return float64(xsect_2.aFull * invLookup(psi, S_Egg[:], N_S_Egg))
	case HORSESHOE:
		return float64(xsect_2.aFull * invLookup(psi, S_Horseshoe[:], N_S_Horseshoe))
	case GOTHIC:
		return float64(xsect_2.aFull * invLookup(psi, S_Gothic[:], N_S_Gothic))
	case CATENARY:
		return float64(xsect_2.aFull * invLookup(psi, S_Catenary[:], N_S_Catenary))
	case SEMIELLIPTICAL:
		return float64(xsect_2.aFull * invLookup(psi, S_SemiEllip[:], N_S_SemiEllip))
	case BASKETHANDLE:
		return float64(xsect_2.aFull * invLookup(psi, S_BasketHandle[:], N_S_BasketHandle))
	case SEMICIRCULAR:
		return float64(xsect_2.aFull * invLookup(psi, S_SemiCirc[:], N_S_SemiCirc))
	default:
		return generic_getAofS(xsect_2, s)
	}
	return 0
}

// xsect.c:1194
func xsect_getdSdA(xsect_2 *TXsect, a_2 float64) float64 {
	switch xsect_2.type_ {
	case FORCE_MAIN, CIRCULAR:
		return circ_getdSdA(xsect_2, a_2)
	case EGGSHAPED:
		return tabular_getdSdA(xsect_2, a_2, S_Egg[:], N_S_Egg)
	case HORSESHOE:
		return tabular_getdSdA(xsect_2, a_2, S_Horseshoe[:], N_S_Horseshoe)
	case GOTHIC:
		return tabular_getdSdA(xsect_2, a_2, S_Gothic[:], N_S_Gothic)
	case CATENARY:
		return tabular_getdSdA(xsect_2, a_2, S_Catenary[:], N_S_Catenary)
	case SEMIELLIPTICAL:
		return tabular_getdSdA(xsect_2, a_2, S_SemiEllip[:], N_S_SemiEllip)
	case BASKETHANDLE:
		return tabular_getdSdA(xsect_2, a_2, S_BasketHandle[:], N_S_BasketHandle)
	case SEMICIRCULAR:
		return tabular_getdSdA(xsect_2, a_2, S_SemiCirc[:], N_S_SemiCirc)
	case RECT_CLOSED:
		return rect_closed_getdSdA(xsect_2, a_2)
	case RECT_OPEN:
		return rect_open_getdSdA(xsect_2, a_2)
	case RECT_TRIANG:
		return rect_triang_getdSdA(xsect_2, a_2)
	case RECT_ROUND:
		return rect_round_getdSdA(xsect_2, a_2)
	case MOD_BASKET:
		return mod_basket_getdSdA(xsect_2, a_2)
	case TRAPEZOIDAL:
		return trapez_getdSdA(xsect_2, a_2)
	case TRIANGULAR:
		return triang_getdSdA(xsect_2, a_2)
	default:
		return generic_getdSdA(xsect_2, a_2)
	}
	return 0
}

// xsect.c:1257
func xsect_getYcrit(xsect_2 *TXsect, q float64) float64 {
	var (
		q2g float64
		y_2 float64
		r   float64
	)
	_, _, _ = q2g, y_2, r
	q2g = float64((float64((q) * (q))) / float64(32.2))
	y_2 = 0
	r = 0
	if q2g == float64(0.0) {
		return float64(0.0)
	}
	switch xsect_2.type_ {
	case DUMMY:
		return float64(0.0)
	case RECT_OPEN, RECT_CLOSED:
		y_2 = M.Pow(float64(q2g/(float64((xsect_2.wMax)*(xsect_2.wMax)))), float64(float64(1.0)/float64(3.0)))
		break
	case TRIANGULAR:
		y_2 = M.Pow(float64(float64(float64(2.0)*q2g)/(float64((xsect_2.sBot)*(xsect_2.sBot)))), float64(float64(1.0)/float64(5.0)))
		break
	case PARABOLIC:
		y_2 = M.Pow(float64(float64(float64(float64(27.0)/float64(32.0))*q2g)/(float64((xsect_2.rBot)*(xsect_2.rBot)))), float64(float64(1.0)/float64(4.0)))
		break
	case POWERFUNC:
		y_2 = float64(float64(1.0) / (float64(float64(float64(2.0)*xsect_2.sBot) + float64(3.0))))
		y_2 = M.Pow(float64(float64(q2g*(float64(xsect_2.sBot+float64(1.0))))/(float64((xsect_2.rBot)*(xsect_2.rBot)))), y_2)
		break
	default:
		y_2 = float64(float64(1.01) * M.Pow(float64(q2g/xsect_2.yFull), float64(float64(1.0)/float64(4.0))))
		if y_2 >= xsect_2.yFull {
			y_2 = float64(float64(0.97) * xsect_2.yFull)
		}
		r = float64(xsect_2.aFull / (float64(float64(float64(3.141592654)/float64(4.0)) * (float64((xsect_2.yFull) * (xsect_2.yFull))))))
		if (r >= float64(0.5)) && (r <= float64(2.0)) {
			y_2 = getYcritEnum(xsect_2, q, y_2)
		} else {
			y_2 = getYcritRidder(xsect_2, q, y_2)
		}
	}
	if (y_2) <= (xsect_2.yFull) {
		return (y_2)
	}
	return (xsect_2.yFull)
}

// xsect.c:1323
func getTransectParams(xsect_2 *TXsect, transect *TTransect) {
	var (
		i    int32
		iMax int32
		wMax float64
		wTbl []float64
	)
	_, _, _, _ = i, iMax, wMax, wTbl
	i = 0
	iMax = 0
	wMax = 0
	wTbl = transect.widthTbl[:]
	xsect_2.yFull = transect.yFull
	xsect_2.wMax = transect.wMax
	xsect_2.aFull = transect.aFull
	xsect_2.rFull = transect.rFull
	xsect_2.sFull = float64(xsect_2.aFull * M.Pow(xsect_2.rFull, float64(float64(2.0)/float64(3.0))))
	xsect_2.sMax = transect.sMax
	xsect_2.aBot = transect.aMax
	iMax = 0
	wMax = wTbl[0]
	for i = 1; i < transect.nTbl; i++ {
		if wTbl[i] < wMax {
			break
		}
		wMax = wTbl[i]
		iMax = i
	}
	xsect_2.ywMax = float64(float64(xsect_2.yFull*float64(iMax)) / (float64(float64((transect.nTbl)) - float64(1))))
}

// xsect.c:1359
func generic_getAofS(xsect_2 *TXsect, s float64) float64 {
	var (
		a_2       float64
		a1        float64
		a2        float64
		tol       float64
		xsectStar TXsectStar
	)
	_, _, _, _, _ = a_2, a1, a2, tol, xsectStar
	a_2 = 0
	a1 = 0
	a2 = 0
	tol = 0
	xsectStar = (TXsectStar{})
	if s <= float64(0.0) {
		return float64(0.0)
	}
	if ((s <= xsect_2.sMax) && (s >= xsect_2.sFull)) && (xsect_2.sMax != xsect_2.sFull) {
		a1 = xsect_2.aFull
		a2 = xsect_getAmax(xsect_2)
	} else {
		a1 = float64(0.0)
		a2 = xsect_getAmax(xsect_2)
	}
	xsectStar.xsect = xsect_2
	xsectStar.s = s
	a_2 = float64(float64(0.5) * (float64(a1 + a2)))
	tol = float64(float64(0.0001) * xsect_2.aFull)
	findroot_Newton(a1, a2, &a_2, tol, evalSofA, &xsectStar)
	return a_2
}

// xsect.c:1404
func evalSofA(a_2 float64, f *float64, df *float64, p any) {
	var (
		xsectStar *TXsectStar
		s         float64
	)
	_, _ = xsectStar, s
	xsectStar = nil
	s = 0
	xsectStar = p.(*TXsectStar)
	s = xsect_getSofA(xsectStar.xsect, a_2)
	*f = float64(s - xsectStar.s)
	*df = xsect_getdSdA(xsectStar.xsect, a_2)
}

// xsect.c:1424
func tabular_getdSdA(xsect_2 *TXsect, a_2 float64, table []float64, nItems int32) float64 {
	var (
		i     int32
		alpha float64
		delta float64
		dSdA  float64
	)
	_, _, _, _ = i, alpha, delta, dSdA
	i = 0
	alpha = float64(a_2 / xsect_2.aFull)
	delta = float64(float64(1.0) / (float64(float64(nItems) - float64(1))))
	dSdA = 0
	i = R.D2I((float64(alpha / delta)))
	if i >= (nItems - 1) {
		i = (nItems - 2)
	}
	dSdA = float64((float64(table[(i+1)] - table[i])) / delta)
	return float64(float64(dSdA*xsect_2.sFull) / xsect_2.aFull)
}

// xsect.c:1453
func generic_getdSdA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		a1     float64
		a2     float64
		alpha  float64
		alpha1 float64
		alpha2 float64
	)
	_, _, _, _, _ = a1, a2, alpha, alpha1, alpha2
	a1 = 0
	a2 = 0
	alpha = float64(a_2 / xsect_2.aFull)
	alpha1 = float64(alpha - float64(0.001))
	alpha2 = float64(alpha + float64(0.001))
	if alpha1 < float64(0.0) {
		alpha1 = float64(0.0)
	}
	a1 = float64(alpha1 * xsect_2.aFull)
	a2 = float64(alpha2 * xsect_2.aFull)
	return float64((float64(xsect_getSofA(xsect_2, a2) - xsect_getSofA(xsect_2, a1))) / (float64(a2 - a1)))
}

// xsect.c:1474
func lookup(x float64, table []float64, nItems int32) float64 {
	var (
		delta float64
		x0    float64
		x1    float64
		y_2   float64
		y2    float64
		i     int32
	)
	_, _, _, _, _, _ = delta, x0, x1, y_2, y2, i
	delta = 0
	x0 = 0
	x1 = 0
	y_2 = 0
	y2 = 0
	i = 0
	delta = float64(float64(1.0) / (float64(float64(nItems) - float64(1))))
	i = R.D2I((float64(x / delta)))
	if i >= (nItems - 1) {
		return table[(nItems - 1)]
	}
	x0 = float64(float64(i) * delta)
	x1 = float64((float64(float64(i) + float64(1))) * delta)
	y_2 = float64(table[i] + float64(float64((float64(x-x0))*(float64(table[(i+1)]-table[i])))/delta))
	if i < 2 {
		y2 = float64(y_2 + float64(float64(float64((float64(x-x0))*(float64(x-x1)))/(float64(delta*delta)))*(float64(float64(float64(table[i]/float64(2.0))-table[(i+1)])+float64(table[(i+2)]/float64(2.0))))))
		if y2 > float64(0.0) {
			y_2 = y2
		}
	}
	if y_2 < float64(0.0) {
		y_2 = float64(0.0)
	}
	return y_2
}

// xsect.c:1511
func invLookup(y_2 float64, table []float64, nItems int32) float64 {
	var (
		dx  float64
		x   float64
		x0  float64
		dy  float64
		n_2 int32
		i   int32
	)
	_, _, _, _, _, _ = dx, x, x0, dy, n_2, i
	dx = 0
	x = 0
	x0 = 0
	dy = 0
	n_2 = 0
	i = 0
	dx = float64(float64(1.0) / (float64(float64(nItems) - float64(1))))
	n_2 = nItems
	if table[(n_2-3)] > table[(n_2-1)] {
		n_2 = (n_2 - 2)
	}
	if (n_2 < nItems) && (y_2 > table[(nItems-1)]) {
		if y_2 >= table[(nItems-3)] {
			return float64((float64(float64(n_2) - float64(1))) * dx)
		}
		if y_2 <= table[(nItems-2)] {
			i = (nItems - 2)
		} else {
			i = (nItems - 3)
		}
	} else {
		i = locate(y_2, table, (n_2 - 1))
	}
	if i >= (n_2 - 1) {
		return float64((float64(float64(n_2) - float64(1))) * dx)
	}
	x0 = float64(float64(i) * dx)
	dy = float64(table[(i+1)] - table[i])
	if dy == float64(0.0) {
		x = x0
	} else {
		x = float64(x0 + float64(float64((float64(y_2-table[i]))*dx)/dy))
	}
	if x < float64(0.0) {
		x = float64(0.0)
	}
	if x > float64(1.0) {
		x = float64(1.0)
	}
	return x
}

// xsect.c:1571
func locate(y_2 float64, table []float64, jLast int32) int32 {
	var (
		j  int32
		j1 int32
		j2 int32
	)
	_, _, _ = j, j1, j2
	j = 0
	j1 = 0
	j2 = jLast
	if y_2 <= table[0] {
		return 0
	}
	if y_2 >= table[jLast] {
		return jLast
	}
	for (j2 - j1) > 1 {
		j = ((j1 + j2) >> 1)
		if y_2 >= table[j] {
			j1 = j
		} else {
			j2 = j
		}
	}
	return j1
}

// xsect.c:1612
func getQcritical(yc float64, p any) float64 {
	var (
		a_2       float64
		w         float64
		qc        float64
		xsectStar *TXsectStar
	)
	_, _, _, _ = a_2, w, qc, xsectStar
	a_2 = 0
	w = 0
	qc = 0
	xsectStar = nil
	xsectStar = p.(*TXsectStar)
	a_2 = xsect_getAofY(xsectStar.xsect, yc)
	w = xsect_getWofY(xsectStar.xsect, yc)
	qc = (-xsectStar.qc)
	if w > float64(0.0) {
		qc = float64(float64(a_2*M.Sqrt(float64(float64(float64(32.2)*a_2)/w))) - xsectStar.qc)
	}
	return qc
}

// xsect.c:1634
func getYcritEnum(xsect_2 *TXsect, q float64, y0 float64) float64 {
	var (
		q0        float64
		dy        float64
		qc        float64
		yc        float64
		i1        int32
		i         int32
		xsectStar TXsectStar
	)
	_, _, _, _, _, _, _ = q0, dy, qc, yc, i1, i, xsectStar
	q0 = 0
	dy = 0
	qc = 0
	yc = 0
	i1 = 0
	i = 0
	xsectStar = (TXsectStar{})
	dy = float64(xsect_2.yFull / float64(25.0))
	i1 = R.D2I((float64(y0 / dy)))
	xsectStar.xsect = xsect_2
	xsectStar.qc = float64(0.0)
	q0 = getQcritical(float64(float64(i1)*dy), &xsectStar)
	if q0 < q {
		yc = xsect_2.yFull
		for i = (i1 + 1); i <= 25; i++ {
			qc = getQcritical(float64(float64(i)*dy), &xsectStar)
			if qc >= q {
				yc = float64((float64(float64((float64(q-q0))/(float64(qc-q0))) + (float64(float64(i) - float64(1))))) * dy)
				break
			}
			q0 = qc
		}
	} else {
		yc = float64(0.0)
		for i = (i1 - 1); i >= 0; i-- {
			qc = getQcritical(float64(float64(i)*dy), &xsectStar)
			if qc < q {
				yc = float64((float64(float64((float64(q-qc))/(float64(q0-qc))) + float64(i))) * dy)
				break
			}
			q0 = qc
		}
	}
	return yc
}

// xsect.c:1700
func getYcritRidder(xsect_2 *TXsect, q float64, y0 float64) float64 {
	var (
		y1        float64
		y2        float64
		yc        float64
		q0        float64
		q1        float64
		q2        float64
		xsectStar TXsectStar
	)
	_, _, _, _, _, _, _ = y1, y2, yc, q0, q1, q2, xsectStar
	y1 = float64(0.0)
	y2 = float64(float64(0.99) * xsect_2.yFull)
	yc = 0
	q0 = 0
	q1 = 0
	q2 = 0
	xsectStar = (TXsectStar{})
	xsectStar.xsect = xsect_2
	xsectStar.qc = float64(0.0)
	q2 = getQcritical(y2, &xsectStar)
	if q2 < q {
		return xsect_2.yFull
	}
	q0 = getQcritical(y0, &xsectStar)
	q1 = getQcritical(float64(float64(0.5)*xsect_2.yFull), &xsectStar)
	if q0 > q {
		y2 = y0
		if q1 < q {
			y1 = float64(float64(0.5) * xsect_2.yFull)
		}
	} else {
		y1 = y0
		if q1 > q {
			y2 = float64(float64(0.5) * xsect_2.yFull)
		}
	}
	xsectStar.qc = q
	yc = findroot_Ridder(y1, y2, float64(0.001), getQcritical, &xsectStar)
	return yc
}

// xsect.c:1755
func rect_closed_getSofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		alfMax float64
	)
	_ = alfMax
	alfMax = float64(0.97)
	if float64(a_2/xsect_2.aFull) > alfMax {
		return float64(xsect_2.sMax + float64(float64((float64(xsect_2.sFull-xsect_2.sMax))*(float64(float64(a_2/xsect_2.aFull)-alfMax)))/(float64(float64(1.0)-alfMax))))
	}
	return float64(a_2 * M.Pow(xsect_getRofA(xsect_2, a_2), float64(float64(2.0)/float64(3.0))))
}

// xsect.c:1770
func rect_closed_getdSdA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		alpha  float64
		alfMax float64
		r      float64
	)
	_, _, _ = alpha, alfMax, r
	alpha = 0
	alfMax = 0
	r = 0
	alfMax = float64(0.97)
	alpha = float64(a_2 / xsect_2.aFull)
	if alpha > alfMax {
		return float64((float64(xsect_2.sFull - xsect_2.sMax)) / (float64((float64(float64(1.0) - alfMax)) * xsect_2.aFull)))
	}
	if alpha <= float64(1.0e-30) {
		return generic_getdSdA(xsect_2, a_2)
	}
	r = xsect_getRofA(xsect_2, a_2)
	return float64((float64(float64(float64(5.0)/float64(3.0)) - float64(float64((float64(float64(2.0)/float64(3.0)))*(float64(float64(2.0)/xsect_2.wMax)))*r))) * M.Pow(r, float64(float64(2.0)/float64(3.0))))
}

// xsect.c:1793
func rect_closed_getRofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		p float64
	)
	_ = p
	p = 0
	if a_2 <= float64(0.0) {
		return float64(0.0)
	}
	p = float64(xsect_2.wMax + float64(float64(float64(2.0)*a_2)/xsect_2.wMax))
	if float64(a_2/xsect_2.aFull) > float64(0.97) {
		p = float64(p + (float64(float64((float64(float64(a_2/xsect_2.aFull)-float64(0.97)))/(float64(float64(1.0)-float64(0.97)))) * xsect_2.wMax)))
	}
	return float64(a_2 / p)
}

// xsect.c:1810
func rect_open_getSofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		y_2 float64
		r   float64
	)
	_, _ = y_2, r
	y_2 = float64(a_2 / xsect_2.wMax)
	r = float64(a_2 / (float64(float64((float64(float64(2.0)-xsect_2.sBot))*y_2) + xsect_2.wMax)))
	return float64(a_2 * M.Pow(r, float64(float64(2.0)/float64(3.0))))
}

// xsect.c:1818
func rect_open_getdSdA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		r    float64
		dPdA float64
	)
	_, _ = r, dPdA
	r = 0
	dPdA = 0
	if float64(a_2/xsect_2.aFull) <= float64(1.0e-30) {
		return generic_getdSdA(xsect_2, a_2)
	}
	r = xsect_getRofA(xsect_2, a_2)
	dPdA = float64((float64(float64(2.0) - xsect_2.sBot)) / xsect_2.wMax)
	return float64((float64(float64(float64(5.0)/float64(3.0)) - float64(float64((float64(float64(2.0)/float64(3.0)))*dPdA)*r))) * M.Pow(r, float64(float64(2.0)/float64(3.0))))
}

// xsect.c:1837
func rect_triang_getYofA(xsect_2 *TXsect, a_2 float64) float64 {
	if a_2 <= xsect_2.aBot {
		return M.Sqrt(float64(a_2 / xsect_2.sBot))
	} else {
		return float64(xsect_2.yBot + float64((float64(a_2-xsect_2.aBot))/xsect_2.wMax))
	}
	return 0
}

// xsect.c:1846
func rect_triang_getRofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		y_2 float64
		p   float64
		alf float64
	)
	_, _, _ = y_2, p, alf
	y_2 = 0
	p = 0
	alf = 0
	if a_2 <= float64(0.0) {
		return float64(0.0)
	}
	y_2 = rect_triang_getYofA(xsect_2, a_2)
	if y_2 <= xsect_2.yBot {
		return float64(a_2 / (float64(float64(float64(2.0)*y_2) * xsect_2.rBot)))
	}
	p = float64(float64(float64(float64(2.0)*xsect_2.yBot)*xsect_2.rBot) + float64(float64(2.0)*(float64(y_2-xsect_2.yBot))))
	alf = float64((float64(a_2 / xsect_2.aFull)) - float64(0.98))
	if alf > float64(0.0) {
		p = float64(p + (float64(float64(alf/(float64(float64(1.0)-float64(0.98)))) * xsect_2.wMax)))
	}
	return float64(a_2 / p)
}

// xsect.c:1866
func rect_triang_getSofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		alfMax float64
	)
	_ = alfMax
	alfMax = float64(0.98)
	if float64(a_2/xsect_2.aFull) > alfMax {
		return float64(xsect_2.sMax + float64(float64((float64(xsect_2.sFull-xsect_2.sMax))*(float64(float64(a_2/xsect_2.aFull)-alfMax)))/(float64(float64(1.0)-alfMax))))
	} else {
		return float64(a_2 * M.Pow(rect_triang_getRofA(xsect_2, a_2), float64(float64(2.0)/float64(3.0))))
	}
	return 0
}

// xsect.c:1879
func rect_triang_getdSdA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		alpha  float64
		alfMax float64
		dPdA   float64
		r      float64
	)
	_, _, _, _ = alpha, alfMax, dPdA, r
	alpha = 0
	alfMax = 0
	dPdA = 0
	r = 0
	alfMax = float64(0.98)
	alpha = float64(a_2 / xsect_2.aFull)
	if alpha > alfMax {
		return float64((float64(xsect_2.sFull - xsect_2.sMax)) / (float64((float64(float64(1.0) - alfMax)) * xsect_2.aFull)))
	}
	if alpha <= float64(1.0e-30) {
		return generic_getdSdA(xsect_2, a_2)
	}
	if a_2 > xsect_2.aBot {
		dPdA = float64(float64(2.0) / xsect_2.wMax)
	} else {
		dPdA = float64(xsect_2.rBot / M.Sqrt(float64(a_2*xsect_2.sBot)))
	}
	r = rect_triang_getRofA(xsect_2, a_2)
	return float64((float64(float64(float64(5.0)/float64(3.0)) - float64(float64((float64(float64(2.0)/float64(3.0)))*dPdA)*r))) * M.Pow(r, float64(float64(2.0)/float64(3.0))))
}

// xsect.c:1902
func rect_triang_getAofY(xsect_2 *TXsect, y_2 float64) float64 {
	if y_2 <= xsect_2.yBot {
		return float64(float64(y_2*y_2) * xsect_2.sBot)
	} else {
		return float64(xsect_2.aBot + float64((float64(y_2-xsect_2.yBot))*xsect_2.wMax))
	}
	return 0
}

// xsect.c:1908
func rect_triang_getRofY(xsect_2 *TXsect, y_2 float64) float64 {
	var (
		p   float64
		a_2 float64
		alf float64
	)
	_, _, _ = p, a_2, alf
	p = 0
	a_2 = 0
	alf = 0
	if y_2 <= xsect_2.yBot {
		return float64(float64(y_2*xsect_2.sBot) / (float64(float64(2.0) * xsect_2.rBot)))
	}
	a_2 = float64(xsect_2.aBot + float64((float64(y_2-xsect_2.yBot))*xsect_2.wMax))
	p = float64(float64(float64(float64(2.0)*xsect_2.yBot)*xsect_2.rBot) + float64(float64(2.0)*(float64(y_2-xsect_2.yBot))))
	alf = float64((float64(a_2 / xsect_2.aFull)) - float64(0.98))
	if alf > float64(0.0) {
		p = float64(p + (float64(float64(alf/(float64(float64(1.0)-float64(0.98)))) * xsect_2.wMax)))
	}
	return float64(a_2 / p)
}

// xsect.c:1927
func rect_triang_getWofY(xsect_2 *TXsect, y_2 float64) float64 {
	if y_2 <= xsect_2.yBot {
		return float64(float64(float64(2.0)*xsect_2.sBot) * y_2)
	} else {
		return xsect_2.wMax
	}
	return 0
}

// xsect.c:1938
func rect_round_getYofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		alpha float64
	)
	_ = alpha
	alpha = 0
	if a_2 > xsect_2.aBot {
		return float64(xsect_2.yBot + float64((float64(a_2-xsect_2.aBot))/xsect_2.wMax))
	}
	alpha = float64(a_2 / (float64(float64(float64(3.141592654)*xsect_2.rBot) * xsect_2.rBot)))
	if alpha < float64(0.04) {
		return float64((float64(float64(2.0) * xsect_2.rBot)) * getYcircular(alpha))
	}
	return float64((float64(float64(2.0) * xsect_2.rBot)) * lookup(alpha, Y_Circ[:], N_Y_Circ))
}

// xsect.c:1952
func rect_round_getRofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		y1     float64
		theta1 float64
		p      float64
		arg    float64
	)
	_, _, _, _ = y1, theta1, p, arg
	y1 = 0
	theta1 = 0
	p = 0
	arg = 0
	if a_2 <= float64(0.0) {
		return float64(0.0)
	}
	if a_2 > xsect_2.aBot {
		y1 = float64((float64(a_2 - xsect_2.aBot)) / xsect_2.wMax)
		theta1 = float64(float64(2.0) * M.Asin(float64(float64(xsect_2.wMax/float64(2.0))/xsect_2.rBot)))
		p = float64(float64(xsect_2.rBot*theta1) + float64(float64(2.0)*y1))
		arg = float64((float64(a_2 / xsect_2.aFull)) - float64(0.98))
		if arg > float64(0.0) {
			p = float64(p + (float64(float64(arg/(float64(float64(1.0)-float64(0.98)))) * xsect_2.wMax)))
		}
		return float64(a_2 / p)
	}
	y1 = rect_round_getYofA(xsect_2, a_2)
	theta1 = float64(float64(2.0) * M.Acos(float64(float64(1.0)-float64(y1/xsect_2.rBot))))
	p = float64(xsect_2.rBot * theta1)
	return float64(a_2 / p)
}

// xsect.c:1978
func rect_round_getSofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		alpha  float64
		aFull  float64
		sFull  float64
		alfMax float64
	)
	_, _, _, _ = alpha, aFull, sFull, alfMax
	alpha = 0
	aFull = 0
	sFull = 0
	alfMax = float64(0.98)
	if float64(a_2/xsect_2.aFull) > alfMax {
		return float64(xsect_2.sMax + float64(float64((float64(xsect_2.sFull-xsect_2.sMax))*(float64(float64(a_2/xsect_2.aFull)-alfMax)))/(float64(float64(1.0)-alfMax))))
	} else if a_2 > xsect_2.aBot {
		return float64(a_2 * M.Pow(xsect_getRofA(xsect_2, a_2), float64(float64(2.0)/float64(3.0))))
	} else {
		aFull = float64(float64(float64(3.141592654)*xsect_2.rBot) * xsect_2.rBot)
		alpha = float64(a_2 / aFull)
		sFull = xsect_2.sBot
		if alpha < float64(0.04) {
			return float64(sFull * getScircular(alpha))
		} else {
			return float64(sFull * lookup(alpha, S_Circ[:], N_S_Circ))
		}
	}
	return 0
}

// xsect.c:2013
func rect_round_getdSdA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		alfMax float64
		r      float64
		dPdA   float64
	)
	_, _, _ = alfMax, r, dPdA
	alfMax = 0
	r = 0
	dPdA = 0
	alfMax = float64(0.98)
	if float64(a_2/xsect_2.aFull) > alfMax {
		return float64((float64(xsect_2.sFull - xsect_2.sMax)) / (float64((float64(float64(1.0) - alfMax)) * xsect_2.aFull)))
	} else if a_2 > xsect_2.aBot {
		r = rect_round_getRofA(xsect_2, a_2)
		dPdA = float64(float64(2.0) / xsect_2.wMax)
		return float64((float64(float64(float64(5.0)/float64(3.0)) - float64(float64((float64(float64(2.0)/float64(3.0)))*dPdA)*r))) * M.Pow(r, float64(float64(2.0)/float64(3.0))))
	} else {
		return generic_getdSdA(xsect_2, a_2)
	}
	return 0
}

// xsect.c:2038
func rect_round_getAofY(xsect_2 *TXsect, y_2 float64) float64 {
	var (
		theta1 float64
	)
	_ = theta1
	theta1 = 0
	if y_2 > xsect_2.yBot {
		return float64(xsect_2.aBot + float64((float64(y_2-xsect_2.yBot))*xsect_2.wMax))
	}
	theta1 = float64(float64(2.0) * M.Acos(float64(float64(1.0)-float64(y_2/xsect_2.rBot))))
	return float64(float64(float64(float64(0.5)*xsect_2.rBot)*xsect_2.rBot) * (float64(theta1 - M.Sin(theta1))))
}

// xsect.c:2051
func rect_round_getRofY(xsect_2 *TXsect, y_2 float64) float64 {
	var (
		theta1 float64
	)
	_ = theta1
	theta1 = 0
	if y_2 <= float64(0.0) {
		return float64(0.0)
	}
	if y_2 > xsect_2.yBot {
		return rect_round_getRofA(xsect_2, rect_round_getAofY(xsect_2, y_2))
	}
	theta1 = float64(float64(2.0) * M.Acos(float64(float64(1.0)-float64(y_2/xsect_2.rBot))))
	return float64(float64(float64(float64(0.5)*xsect_2.rBot)*(float64(float64(1.0)-M.Sin(theta1)))) / theta1)
}

// xsect.c:2065
func rect_round_getWofY(xsect_2 *TXsect, y_2 float64) float64 {
	if y_2 > xsect_2.yBot {
		return xsect_2.wMax
	}
	return float64(float64(2.0) * M.Sqrt(float64(y_2*(float64(float64(float64(2.0)*xsect_2.rBot)-y_2)))))
}

// xsect.c:2082
func mod_basket_getYofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		alpha float64
		y1    float64
	)
	_, _ = alpha, y1
	alpha = 0
	y1 = 0
	if a_2 <= float64(xsect_2.aFull-xsect_2.aBot) {
		return float64(a_2 / xsect_2.wMax)
	}
	alpha = float64((float64(xsect_2.aFull - a_2)) / (float64(float64(float64(3.141592654)*xsect_2.rBot) * xsect_2.rBot)))
	if alpha < float64(0.04) {
		y1 = getYcircular(alpha)
	} else {
		y1 = lookup(alpha, Y_Circ[:], N_Y_Circ)
	}
	y1 = float64(float64(float64(2.0)*xsect_2.rBot) * y1)
	return float64(xsect_2.yFull - y1)
}

// xsect.c:2101
func mod_basket_getRofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		y1     float64
		p      float64
		theta1 float64
	)
	_, _, _ = y1, p, theta1
	y1 = 0
	p = 0
	theta1 = 0
	if a_2 <= float64(xsect_2.aFull-xsect_2.aBot) {
		return float64(a_2 / (float64(xsect_2.wMax + float64(float64(float64(2.0)*a_2)/xsect_2.wMax))))
	}
	y1 = float64(xsect_2.yFull - mod_basket_getYofA(xsect_2, a_2))
	theta1 = float64(float64(2.0) * M.Acos(float64(float64(1.0)-float64(y1/xsect_2.rBot))))
	p = float64((float64(xsect_2.sBot - theta1)) * xsect_2.rBot)
	y1 = float64(xsect_2.yFull - xsect_2.yBot)
	p = float64(float64(p+float64(float64(2.0)*y1)) + xsect_2.wMax)
	return float64(a_2 / p)
}

// xsect.c:2128
func mod_basket_getdSdA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		r    float64
		dPdA float64
	)
	_, _ = r, dPdA
	r = 0
	dPdA = 0
	if (a_2 <= float64(xsect_2.aFull-xsect_2.aBot)) && (float64(a_2/xsect_2.aFull) > float64(1.0e-30)) {
		r = float64(a_2 / (float64(xsect_2.wMax + float64(float64(float64(2.0)*a_2)/xsect_2.wMax))))
		dPdA = float64(float64(2.0) / xsect_2.wMax)
		return float64((float64(float64(float64(5.0)/float64(3.0)) - float64(float64((float64(float64(2.0)/float64(3.0)))*dPdA)*r))) * M.Pow(r, float64(float64(2.0)/float64(3.0))))
	} else {
		return generic_getdSdA(xsect_2, a_2)
	}
	return 0
}

// xsect.c:2145
func mod_basket_getAofY(xsect_2 *TXsect, y_2 float64) float64 {
	var (
		a1     float64
		theta1 float64
		y1     float64
	)
	_, _, _ = a1, theta1, y1
	a1 = 0
	theta1 = 0
	y1 = 0
	if y_2 <= float64(xsect_2.yFull-xsect_2.yBot) {
		return float64(y_2 * xsect_2.wMax)
	}
	y1 = float64(xsect_2.yFull - y_2)
	theta1 = float64(float64(2.0) * M.Acos(float64(float64(1.0)-float64(y1/xsect_2.rBot))))
	a1 = float64(float64(float64(float64(0.5)*xsect_2.rBot)*xsect_2.rBot) * (float64(theta1 - M.Sin(theta1))))
	return float64(xsect_2.aFull - a1)
}

// xsect.c:2162
func mod_basket_getWofY(xsect_2 *TXsect, y_2 float64) float64 {
	var (
		y1 float64
	)
	_ = y1
	y1 = 0
	if y_2 <= float64(0.0) {
		return float64(0.0)
	}
	if y_2 <= float64(xsect_2.yFull-xsect_2.yBot) {
		return xsect_2.wMax
	}
	y1 = float64(xsect_2.yFull - y_2)
	return float64(float64(2.0) * M.Sqrt(float64(y1*(float64(float64(float64(2.0)*xsect_2.rBot)-y1)))))
}

// xsect.c:2184
func trapez_getYofA(xsect_2 *TXsect, a_2 float64) float64 {
	if xsect_2.sBot == float64(0.0) {
		return float64(a_2 / xsect_2.yBot)
	}
	return float64((float64(M.Sqrt(float64(float64(xsect_2.yBot*xsect_2.yBot)+float64(float64(float64(4.0)*xsect_2.sBot)*a_2))) - xsect_2.yBot)) / (float64(float64(2.0) * xsect_2.sBot)))
}

// xsect.c:2191
func trapez_getRofA(xsect_2 *TXsect, a_2 float64) float64 {
	return float64(a_2 / (float64(xsect_2.yBot + float64(trapez_getYofA(xsect_2, a_2)*xsect_2.rBot))))
}

// xsect.c:2196
func trapez_getdSdA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		r    float64
		dPdA float64
	)
	_, _ = r, dPdA
	r = 0
	dPdA = 0
	if float64(a_2/xsect_2.aFull) <= float64(1.0e-30) {
		return generic_getdSdA(xsect_2, a_2)
	}
	r = trapez_getRofA(xsect_2, a_2)
	dPdA = float64(xsect_2.rBot / M.Sqrt(float64(float64(xsect_2.yBot*xsect_2.yBot)+float64(float64(float64(4.0)*xsect_2.sBot)*a_2))))
	return float64((float64(float64(float64(5.0)/float64(3.0)) - float64(float64((float64(float64(2.0)/float64(3.0)))*dPdA)*r))) * M.Pow(r, float64(float64(2.0)/float64(3.0))))
}

// xsect.c:2210
func trapez_getAofY(xsect_2 *TXsect, y_2 float64) float64 {
	return float64((float64(xsect_2.yBot + float64(xsect_2.sBot*y_2))) * y_2)
}

// xsect.c:2215
func trapez_getRofY(xsect_2 *TXsect, y_2 float64) float64 {
	if y_2 == float64(0.0) {
		return float64(0.0)
	}
	return float64(trapez_getAofY(xsect_2, y_2) / (float64(xsect_2.yBot + float64(y_2*xsect_2.rBot))))
}

// xsect.c:2221
func trapez_getWofY(xsect_2 *TXsect, y_2 float64) float64 {
	return float64(xsect_2.yBot + float64(float64(float64(2.0)*y_2)*xsect_2.sBot))
}

// xsect.c:2231
func triang_getYofA(xsect_2 *TXsect, a_2 float64) float64 {
	return M.Sqrt(float64(a_2 / xsect_2.sBot))
}

// xsect.c:2236
func triang_getRofA(xsect_2 *TXsect, a_2 float64) float64 {
	return float64(a_2 / (float64(float64(float64(2.0)*triang_getYofA(xsect_2, a_2)) * xsect_2.rBot)))
}

// xsect.c:2241
func triang_getdSdA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		r    float64
		dPdA float64
	)
	_, _ = r, dPdA
	r = 0
	dPdA = 0
	if float64(a_2/xsect_2.aFull) <= float64(1.0e-30) {
		return generic_getdSdA(xsect_2, a_2)
	}
	r = triang_getRofA(xsect_2, a_2)
	dPdA = float64(xsect_2.rBot / M.Sqrt(float64(a_2*xsect_2.sBot)))
	return float64((float64(float64(float64(5.0)/float64(3.0)) - float64(float64((float64(float64(2.0)/float64(3.0)))*dPdA)*r))) * M.Pow(r, float64(float64(2.0)/float64(3.0))))
}

// xsect.c:2253
func triang_getAofY(xsect_2 *TXsect, y_2 float64) float64 {
	return float64(float64(y_2*y_2) * xsect_2.sBot)
}

// xsect.c:2258
func triang_getRofY(xsect_2 *TXsect, y_2 float64) float64 {
	return float64((float64(y_2 * xsect_2.sBot)) / (float64(float64(2.0) * xsect_2.rBot)))
}

// xsect.c:2263
func triang_getWofY(xsect_2 *TXsect, y_2 float64) float64 {
	return float64(float64(float64(2.0)*xsect_2.sBot) * y_2)
}

// xsect.c:2273
func parab_getYofA(xsect_2 *TXsect, a_2 float64) float64 {
	return M.Pow(float64(float64((float64(float64(3.0)/float64(4.0)))*a_2)/xsect_2.rBot), float64(float64(2.0)/float64(3.0)))
}

// xsect.c:2278
func parab_getRofA(xsect_2 *TXsect, a_2 float64) float64 {
	if a_2 <= float64(0.0) {
		return float64(0.0)
	}
	return float64(a_2 / parab_getPofY(xsect_2, parab_getYofA(xsect_2, a_2)))
}

// xsect.c:2284
func parab_getPofY(xsect_2 *TXsect, y_2 float64) float64 {
	var (
		x float64
		t float64
	)
	_, _ = x, t
	x = float64(float64(float64(2.0)*M.Sqrt(y_2)) / xsect_2.rBot)
	t = M.Sqrt(float64(float64(1.0) + float64(x*x)))
	return float64(float64(float64(float64(0.5)*xsect_2.rBot)*xsect_2.rBot) * (float64(float64(x*t) + M.Log(float64(x+t)))))
}

// xsect.c:2291
func parab_getAofY(xsect_2 *TXsect, y_2 float64) float64 {
	return (float64(float64(float64(float64(float64(4.0)/float64(3.0))*xsect_2.rBot)*y_2) * M.Sqrt(y_2)))
}

// xsect.c:2296
func parab_getRofY(xsect_2 *TXsect, y_2 float64) float64 {
	if y_2 <= float64(0.0) {
		return float64(0.0)
	}
	return float64(parab_getAofY(xsect_2, y_2) / parab_getPofY(xsect_2, y_2))
}

// xsect.c:2302
func parab_getWofY(xsect_2 *TXsect, y_2 float64) float64 {
	return float64(float64(float64(2.0)*xsect_2.rBot) * M.Sqrt(y_2))
}

// xsect.c:2312
func powerfunc_getYofA(xsect_2 *TXsect, a_2 float64) float64 {
	return M.Pow(float64(a_2/xsect_2.rBot), float64(float64(1.0)/(float64(xsect_2.sBot+float64(1.0)))))
}

// xsect.c:2317
func powerfunc_getRofA(xsect_2 *TXsect, a_2 float64) float64 {
	if a_2 <= float64(0.0) {
		return float64(0.0)
	}
	return float64(a_2 / powerfunc_getPofY(xsect_2, powerfunc_getYofA(xsect_2, a_2)))
}

// xsect.c:2323
func powerfunc_getPofY(xsect_2 *TXsect, y_2 float64) float64 {
	var (
		dy1 float64
		h   float64
		m   float64
		p   float64
		y1  float64
		x1  float64
		x2  float64
		y2  float64
		dx  float64
		dy  float64
	)
	_, _, _, _, _, _, _, _, _, _ = dy1, h, m, p, y1, x1, x2, y2, dx, dy
	dy1 = float64(float64(0.02) * xsect_2.yFull)
	h = float64(float64((float64(xsect_2.sBot+float64(1.0)))*xsect_2.rBot) / float64(2.0))
	m = xsect_2.sBot
	p = float64(0.0)
	y1 = float64(0.0)
	x1 = float64(0.0)
	x2 = 0
	y2 = 0
	dx = 0
	dy = 0
	for __dw1 := true; __dw1 || (y2 < y_2); __dw1 = false {
		y2 = float64(y1 + dy1)
		if y2 > y_2 {
			y2 = y_2
		}
		x2 = float64(h * M.Pow(y2, m))
		dx = float64(x2 - x1)
		dy = float64(y2 - y1)
		p = float64(p + (M.Sqrt(float64(float64(dx*dx) + float64(dy*dy)))))
		x1 = x2
		y1 = y2
	}
	return float64(float64(2.0) * p)
}

// xsect.c:2346
func powerfunc_getAofY(xsect_2 *TXsect, y_2 float64) float64 {
	return float64(xsect_2.rBot * M.Pow(y_2, float64(xsect_2.sBot+float64(1.0))))
}

// xsect.c:2351
func powerfunc_getRofY(xsect_2 *TXsect, y_2 float64) float64 {
	if y_2 <= float64(0.0) {
		return float64(0.0)
	}
	return float64(powerfunc_getAofY(xsect_2, y_2) / powerfunc_getPofY(xsect_2, y_2))
}

// xsect.c:2357
func powerfunc_getWofY(xsect_2 *TXsect, y_2 float64) float64 {
	return float64(float64((float64(xsect_2.sBot+float64(1.0)))*xsect_2.rBot) * M.Pow(y_2, xsect_2.sBot))
}

// xsect.c:2367
func circ_getYofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		alpha float64
	)
	_ = alpha
	alpha = float64(a_2 / xsect_2.aFull)
	if alpha < float64(0.04) {
		return float64(xsect_2.yFull * getYcircular(alpha))
	} else {
		return float64(xsect_2.yFull * lookup(alpha, Y_Circ[:], N_Y_Circ))
	}
	return 0
}

// xsect.c:2378
func circ_getAofS(xsect_2 *TXsect, s float64) float64 {
	var (
		psi float64
	)
	_ = psi
	psi = float64(s / xsect_2.sFull)
	if psi == float64(0.0) {
		return float64(0.0)
	}
	if psi >= float64(1.0) {
		return xsect_2.aFull
	}
	if psi <= float64(0.015) {
		return float64(xsect_2.aFull * getAcircular(psi))
	} else {
		return float64(xsect_2.aFull * invLookup(psi, S_Circ[:], N_S_Circ))
	}
	return 0
}

// xsect.c:2391
func circ_getSofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		alpha float64
	)
	_ = alpha
	alpha = float64(a_2 / xsect_2.aFull)
	if alpha < float64(0.04) {
		return float64(xsect_2.sFull * getScircular(alpha))
	} else {
		return float64(xsect_2.sFull * lookup(alpha, S_Circ[:], N_S_Circ))
	}
	return 0
}

// xsect.c:2403
func circ_getdSdA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		alpha float64
		theta float64
		p     float64
		r     float64
		dPdA  float64
	)
	_, _, _, _, _ = alpha, theta, p, r, dPdA
	alpha = 0
	theta = 0
	p = 0
	r = 0
	dPdA = 0
	alpha = float64(a_2 / xsect_2.aFull)
	if alpha <= float64(1.0e-30) {
		return float64(1.0e-30)
	} else if alpha < float64(0.04) {
		theta = getThetaOfAlpha(alpha)
		p = float64(float64(theta*xsect_2.yFull) / float64(2.0))
		r = float64(a_2 / p)
		dPdA = float64(float64(float64(4.0)/xsect_2.yFull) / (float64(float64(1.0) - M.Cos(theta))))
		return float64((float64(float64(float64(5.0)/float64(3.0)) - float64(float64((float64(float64(2.0)/float64(3.0)))*dPdA)*r))) * M.Pow(r, float64(float64(2.0)/float64(3.0))))
	} else {
		return tabular_getdSdA(xsect_2, a_2, S_Circ[:], N_S_Circ)
	}
	return 0
}

// xsect.c:2450
func circ_getAofY(xsect_2 *TXsect, y_2 float64) float64 {
	var (
		yNorm float64
	)
	_ = yNorm
	yNorm = 0
	yNorm = float64(y_2 / xsect_2.yFull)
	return float64(xsect_2.aFull * lookup(yNorm, A_Circ[:], N_A_Circ))
}

// xsect.c:2462
func filled_circ_getYofA(xsect_2 *TXsect, a_2 float64) float64 {
	var (
		y_2 float64
	)
	_ = y_2
	y_2 = 0
	xsect_2.yFull = float64(xsect_2.yFull + xsect_2.yBot)
	xsect_2.aFull = float64(xsect_2.aFull + xsect_2.aBot)
	a_2 = float64(a_2 + xsect_2.aBot)
	y_2 = circ_getYofA(xsect_2, a_2)
	y_2 = float64(y_2 - xsect_2.yBot)
	xsect_2.yFull = float64(xsect_2.yFull - xsect_2.yBot)
	xsect_2.aFull = float64(xsect_2.aFull - xsect_2.aBot)
	return y_2
}

// xsect.c:2481
func filled_circ_getAofY(xsect_2 *TXsect, y_2 float64) float64 {
	var (
		a_2 float64
	)
	_ = a_2
	a_2 = 0
	xsect_2.yFull = float64(xsect_2.yFull + xsect_2.yBot)
	xsect_2.aFull = float64(xsect_2.aFull + xsect_2.aBot)
	y_2 = float64(y_2 + xsect_2.yBot)
	a_2 = circ_getAofY(xsect_2, y_2)
	a_2 = float64(a_2 - xsect_2.aBot)
	xsect_2.yFull = float64(xsect_2.yFull - xsect_2.yBot)
	xsect_2.aFull = float64(xsect_2.aFull - xsect_2.aBot)
	return a_2
}

// xsect.c:2500
func filled_circ_getRofY(xsect_2 *TXsect, y_2 float64) float64 {
	var (
		a_2 float64
		r   float64
		p   float64
	)
	_, _, _ = a_2, r, p
	a_2 = 0
	r = 0
	p = 0
	xsect_2.yFull = float64(xsect_2.yFull + xsect_2.yBot)
	xsect_2.aFull = float64(xsect_2.aFull + xsect_2.aBot)
	y_2 = float64(y_2 + xsect_2.yBot)
	a_2 = circ_getAofY(xsect_2, y_2)
	r = float64(float64(float64(0.25)*xsect_2.yFull) * lookup(float64(y_2/xsect_2.yFull), R_Circ[:], N_R_Circ))
	p = (float64(a_2 / r))
	a_2 = float64(a_2 - xsect_2.aBot)
	p = float64(float64(p-xsect_2.rBot) + xsect_2.sBot)
	r = float64(a_2 / p)
	xsect_2.yFull = float64(xsect_2.yFull - xsect_2.yBot)
	xsect_2.aFull = float64(xsect_2.aFull - xsect_2.aBot)
	return r
}

// xsect.c:2531
func getYcircular(alpha float64) float64 {
	var (
		theta float64
	)
	_ = theta
	theta = 0
	if alpha >= float64(1.0) {
		return float64(1.0)
	}
	if alpha <= float64(0.0) {
		return float64(0.0)
	}
	if alpha <= float64(1.0e-5) {
		theta = M.Pow(float64(float64(37.6911)*alpha), float64(float64(1.0)/float64(3.0)))
		return float64(float64(theta*theta) / float64(16.0))
	}
	theta = getThetaOfAlpha(alpha)
	return float64((float64(float64(1.0) - M.Cos(float64(theta/float64(2.0))))) / float64(2.0))
}

// xsect.c:2545
func getScircular(alpha float64) float64 {
	var (
		theta float64
	)
	_ = theta
	theta = 0
	if alpha >= float64(1.0) {
		return float64(1.0)
	}
	if alpha <= float64(0.0) {
		return float64(0.0)
	}
	if alpha <= float64(1.0e-5) {
		theta = M.Pow(float64(float64(37.6911)*alpha), float64(float64(1.0)/float64(3.0)))
		return float64(M.Pow(theta, float64(float64(13.0)/float64(3.0))) / float64(124.4797))
	}
	theta = getThetaOfAlpha(alpha)
	return float64(float64(M.Pow((float64(theta-M.Sin(theta))), float64(float64(5.0)/float64(3.0)))/(float64(float64(2.0)*float64(3.141592654)))) / M.Pow(theta, float64(float64(2.0)/float64(3.0))))
}

// xsect.c:2559
func getAcircular(psi float64) float64 {
	var (
		theta float64
	)
	_ = theta
	theta = 0
	if psi >= float64(1.0) {
		return float64(1.0)
	}
	if psi <= float64(0.0) {
		return float64(0.0)
	}
	if psi <= float64(1.0e-6) {
		theta = M.Pow(float64(float64(124.4797)*psi), float64(float64(3.0)/float64(13.0)))
		return float64(float64(float64(theta*theta)*theta) / float64(37.6911))
	}
	theta = getThetaOfPsi(psi)
	return float64((float64(theta - M.Sin(theta))) / (float64(float64(2.0) * float64(3.141592654))))
}

// xsect.c:2573
func getThetaOfAlpha(alpha float64) float64 {
	var (
		k      int32
		theta  float64
		theta1 float64
		ap     float64
		d      float64
	)
	_, _, _, _, _ = k, theta, theta1, ap, d
	k = 0
	theta = 0
	theta1 = 0
	ap = 0
	d = 0
	if alpha > float64(0.04) {
		theta = float64(float64(1.2) + float64(float64(float64(5.08)*(float64(alpha-float64(0.04))))/float64(0.96)))
	} else {
		theta = float64(float64(float64(0.031715)-float64(float64(12.79384)*alpha)) + float64(float64(8.28479)*M.Sqrt(alpha)))
	}
	theta1 = theta
	ap = float64((float64(float64(2.0) * float64(3.141592654))) * alpha)
	for k = 1; k <= 40; k++ {
		d = float64((-(float64(float64(ap-theta) + M.Sin(theta)))) / (float64(float64(1.0) - M.Cos(theta))))
		if d > float64(1.0) {
			if (d) >= float64(0.0) {
				d = M.Fabs(float64(1.0))
			} else {
				d = (-(M.Fabs(float64(1.0))))
			}
		}
		theta = float64(theta - d)
		if M.Fabs(d) <= float64(0.0001) {
			return theta
		}
	}
	return theta1
}

// xsect.c:2593
func getThetaOfPsi(psi float64) float64 {
	var (
		k      int32
		theta  float64
		theta1 float64
		ap     float64
		tt     float64
		tt23   float64
		t3     float64
		d      float64
	)
	_, _, _, _, _, _, _, _ = k, theta, theta1, ap, tt, tt23, t3, d
	k = 0
	theta = 0
	theta1 = 0
	ap = 0
	tt = 0
	tt23 = 0
	t3 = 0
	d = 0
	if psi > float64(0.90) {
		theta = float64(float64(4.17) + float64(float64(float64(1.12)*(float64(psi-float64(0.90))))/float64(0.176)))
	} else if psi > float64(0.5) {
		theta = float64(float64(3.14) + float64(float64(float64(1.03)*(float64(psi-float64(0.5))))/float64(0.4)))
	} else if psi > float64(0.015) {
		theta = float64(float64(1.2) + float64(float64(float64(1.94)*(float64(psi-float64(0.015))))/float64(0.485)))
	} else {
		theta = float64(float64(float64(0.12103)-float64(float64(55.5075)*psi)) + float64(float64(15.62254)*M.Sqrt(psi)))
	}
	theta1 = theta
	ap = float64((float64(float64(2.0) * float64(3.141592654))) * psi)
	for k = 1; k <= 40; k++ {
		theta = M.Fabs(theta)
		tt = float64(theta - M.Sin(theta))
		tt23 = M.Pow(tt, float64(float64(2.0)/float64(3.0)))
		t3 = M.Pow(theta, float64(float64(1.0)/float64(3.0)))
		d = float64(float64(float64(ap*theta)/t3) - float64(tt*tt23))
		d = float64(d / (float64(float64(float64(ap*(float64(float64(2.0)/float64(3.0))))/t3) - float64(float64((float64(float64(5.0)/float64(3.0)))*tt23)*(float64(float64(1.0)-M.Cos(theta)))))))
		theta = float64(theta - d)
		if M.Fabs(d) <= float64(0.0001) {
			return theta
		}
	}
	return theta1
}
