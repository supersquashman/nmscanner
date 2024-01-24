
class ImmTools:
    
    def print_vnum_object_adjustments(self, vnums, level_req):
        for vnum in vnums:
            print(f"oinvoke {vnum}")
            print(f"oset {vnum} flags prototype")
            print(f"oset {vnum} level {level_req}")
            print(f"oset {vnum} legacy_level 100")
            print(f"opedit {vnum} delete 1")
            print(f"oset {vnum} flags prototype")
            print(f"eat {vnum}")
    
    def print_bind_adjustment(self,vnums):
        for vnum in vnums:
            print(f"oinvoke {vnum}")
            print(f"oset {vnum} flags prototype")
            print(f"oset {vnum} flags binding")
            print(f"oset {vnum} flags prototype")
            print(f"eat {vnum}")

commander = ImmTools()

yamayuki_vnums = {1578,1579,1580,1581,1582,1583,1584,1586,1589,1593,1594,1595,1596,1597,
                  1598,1599,1602,1603,1616,1617,1618,1619,2374,2377,2381,2383,2385,
                  2387,2389,2391,2393,2395,2397,2404}

add_binding = {1598, 1599, 1608, 1614, 1604, 1631, 1615,2375,2376,2377,2378,2653,2382,
               2384,2386,2388,2390,2392,2394,2396,2402,2398,2403,2405,2406,2407,2408,
               2409,2410,2411,2412,2413,2414,2415,2416,2417,2418,2776,2966,2968,2970,
               2972,2974,2976,2978,2994,2995,2997,2998,3000,3002,3004,3005,3006,3007}

commander.print_vnum_object_adjustments(yamayuki_vnums, 150)
commander.print_bind_adjustment(add_binding)

hikarikumo_vnums = {2965,2967,2969,2971,2973,2975,2977,2982,2993,2996,2999,3001,3003}
commander.print_vnum_object_adjustments(hikarikumo_vnums, 250)

#add special condition to prevent legaacy from using {1590,,} ?