clear all;
close all;
clc;

system = 0;
betaLowerLimit = 40;
betaUpperLimit = 40;
betaStep = 10;

n0s = 1:100:10000;
n = 5000;
m = 10;

figure;
hold on;
for beta = betaLowerLimit:betaStep:betaUpperLimit
    fname = ['./section2final/beta-', num2str(beta) ,'-system', num2str(system) ,'.csv'];
    data = csvread(fname, 1);

    Sweep = data(:,1);
    Beta = data(:,2);
    SubSystemID = data(:,3);
    Magnetisation = data(:,4);
    DimensionlessEnergy = data(:,5);
    
    AverageMagnetisations = [];
    AverageEnergies = [];
    
    for n0 = n0s
        magnetisations = Magnetisation(n0:n:(n0+n*m))
        energies = DimensionlessEnergy(n0:n:(n0+n*m))
        AverageMagnetisations = [AverageMagnetisations mean(magnetisations)];
        AverageEnergies = [AverageEnergies mean(energies)]
    end
    
    plot(n0s, AverageMagnetisations, 'x-', 'DisplayName', sprintf('$\\beta = %0.2f$', beta/100))
end

hold off;
xlabel('$n_0$', 'Interpreter', 'latex', 'FontSize', 16);
ylabel('$<\mathcal{M}>$', 'Interpreter', 'latex', 'FontSize', 16);
% ylim([-1, 0.2])

legend_handle = legend('-DynamicLegend');
set(legend_handle,'Interpreter','latex')
legend('show');

 
% figure;
% plot(Sweep, DimensionlessEnergy)
% ylabel('$E/J$', 'Interpreter', 'latex', 'FontSize', 16);