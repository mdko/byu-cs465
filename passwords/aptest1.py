import argparse
parser = argparse.ArgumentParser()
parser.add_argument("square", help="echo the string you use here", type=int)
args = parser.parse_args()
print args.square**2
