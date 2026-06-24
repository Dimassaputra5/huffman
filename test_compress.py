import subprocess, pathlib, sys

ROOT = pathlib.Path(__file__).parent
INPUT = ROOT / "input"
OUT = ROOT / "out"
EXE = ROOT / "huffman.exe"

def run(cmd):
    r = subprocess.run(cmd, capture_output=True, text=True)
    if r.returncode:
        print(r.stderr, end="")
        sys.exit(1)
    return r.stdout

print("Compressing...")
files = sorted(INPUT.iterdir())
for f in files:
    stem = f.stem
    huf = OUT / f"{stem}.huf"
    run([str(EXE), "compress", str(f), "-o", str(huf), "-f"])

print("Decompressing...")
dec = OUT / "verify"
dec.mkdir(exist_ok=True)
for f in files:
    stem = f.stem
    huf = OUT / f"{stem}.huf"
    out = dec / f.name
    run([str(EXE), "decompress", str(huf), "-o", str(out), "-f"])

print("Validating...")
print()
print(f"{'File':<15} {'Original':>10} {'Compressed':>12} {'Ratio':>7}  {'Status'}")
print(f"{'----':<15} {'--------':>10} {'----------':>12} {'-----':>7}  {'------'}")
for f in files:
    stem = f.stem
    orig = f.read_bytes()
    comp = (OUT / f"{stem}.huf").read_bytes()
    decomp = (dec / f.name).read_bytes()
    ok = orig == decomp
    ratio = (len(comp) / len(orig)) * 100 if orig else 0
    print(f"{f.name:<15} {len(orig):>10} {len(comp):>12} {ratio:>6.1f}%  {'PASS' if ok else 'FAIL'}")
    if not ok:
        sys.exit(1)

print()
print("All passed.")
