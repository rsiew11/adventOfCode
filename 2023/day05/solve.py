
def read(f):
    with open(f, "r+") as fd:
        return list(
           fd.read().splitlines() 
        )

def main():
    data = read("./input.txt")
    seeds = list(map(int, data[0].split("seeds:")[1].strip().split()))
    data = data[2:]

    # per map
    # dst, src, rng
    # (dst --> dst + rng), (src --> src + rng)
    map_names = ["seed_soil", "soil_fert", "fert_watr", "watr_light", "lght_temp", "temp_humi", "humi_loc"]
    maps = [[] for _ in range(len(map_names))]


if __name__ == "__main__":
    main()